package initialize

import (
	"errors"
	"fmt"
	"honor_money/cmd"
	"honor_money/def"
	"honor_money/glade"
	"honor_money/tool"
	"sync"
	"time"
)
var mutex sync.Mutex
//全局chan控制定时器开关
//当任务开始执行定期检测任务
func StartTask(dev *cmd.AdbHandler) {
	mutex.Lock()
	dev.IsEnd =false
	//开始任务
	//初始化各种图标
	err :=InitImage(dev)
	if err!=nil{
		glade.MsgError(def.AppWin1,err.Error())
		return
	}
	go func(dev *cmd.AdbHandler) {
		CStep.Step=Start
		CStep.WaitTime=time.Now()
		err:=PointStep(dev)
		if err!=nil{
			StopTask(dev,STOP)
			glade.NoticeError(err.Error())
		}else{
			fmt.Println("任务定位停止且没有错误")
		}

	}(dev)
	mutex.Unlock()
}
type Tag uint8
const (
	CHANGE Tag = iota //应用当前任务执行完毕请求切换任务
	PAUSE //用户暂停任务
	STOP //所有任务完成并执行结束任务的操作，关应用，关机等

)
//暂停所有任务
func StopTask(dev *cmd.AdbHandler,tag Tag) {
	mutex.Lock()

	fmt.Println("停止任务")
	for index,ch:=range dev.Timers{//关闭所有定时器
		if ch!=nil{
			fmt.Println("当前停止chan",index,ch)
			dev.Timers[index] <- true
			close(dev.Timers[index])
			dev.Timers[index] = nil
		}
	}

	if tag == CHANGE{
		//执行切换任务，后续开发
	}else if tag == PAUSE{
		dev.IsEnd =true
		_,err := dev.Cmd(" shell wm size reset ")
		if err!=nil{
		    fmt.Printf("shell wm size reset failed,err:%v",err)
		    return
		}
		fmt.Println("执行暂停任务")
		def.Button1.SetLabel("开始")
		def.Entry1.SetSensitive(true)
	}else if tag == STOP{
		dev.IsEnd =true
		_,err := dev.Cmd(" shell wm size reset ")
		if err!=nil{
			fmt.Printf("shell wm size reset failed,err:%v",err)
			return
		}
		EndHandler(dev)//执行结束操作
	}

	mutex.Unlock()
}
func EndHandler(dev *cmd.AdbHandler) {

	//判断是否需要关机和关闭应用
	isCheck := def.Check1.GetActive()
	if isCheck {
		dev.Cmd("shell am force-stop " + def.GamePackName)
	}
}
//暂时单设备就这里写固定
type CommonStep struct {
	Step int
	WaitTime time.Time
}
var CStep = &CommonStep{
	Step:     Start,
	WaitTime: time.Now(),
}
const (
	Start int = iota //需要启动应用
	ALL //全部判断
	Login //需要点击登录

	//CLOSE //判断是否有遮罩需要点叉叉
	//RETURN //需要返回到首页
	INDEX //进入首页 可能需要 点击返回 和 点close 或者服务器断开连接 这里单独函数处理

	//后面操作一起操作，长时间不能完成操作就重启应用

)


//PointStep 定位当前需要的操作
func PointStep(dev *cmd.AdbHandler)(err error){
	CStep.WaitTime=time.Now()
	devImgs:=Imgs[dev.Size.String()]
	CStep.Step = Start
	for{
		if dev.IsEnd {
			return nil
		}
		if time.Since(CStep.WaitTime) >120*time.Second{//如果每个流程停留超过60秒就停止所有任务并提示错误
			return errors.New("流程等待时间超时60s")
		}
		if CStep.Step == Start{//这里不用判断all，启用应用会将应用放到前台
			dev.Cmd(" shell am start -n "+def.GamePackFullName)
			fmt.Println("启动应用")
			CStep.Step = ALL //进入登录判断
		}
		img,err:=glade.PullPicture(dev)//获取当前设备图片

		if err!=nil{
			glade.NoticeError(err.Error())
		}

		//登录判断
		if CStep.Step==Login||CStep.Step==ALL{
			fmt.Println("登录判断")
			//对比登录图片

			//超时判断
			ok,p,err:=tool.MatchPicture(devImgs["login"],img)
			if err!=nil{
				glade.NoticeError("登录图片匹配错误:"+err.Error())
			    return err
			}
			if ok{
				dev.Tap(p.X,p.Y)
				CStep.Step=INDEX
				continue
			}

		}

		if CStep.Step == INDEX  || CStep.Step==ALL {
			//首页判断，这里是终点

			ok,err:=CheckIndex(dev,img)
			if err!=nil{
				return err
			}
			if ok{
				return nil
			}
			//遮罩判断 首页判断 返回判断 服务器断开判断
			ok,err=CheckMain(dev,img)
			if err!=nil{
				return errors.New(err.Error())
			}
			if ok{
				CStep.Step=INDEX
			}else{
				continue
			}

		}


		time.Sleep(2*time.Second)//睡一秒，避免太频繁的
	}

}


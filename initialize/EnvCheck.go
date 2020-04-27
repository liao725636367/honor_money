package initialize

import (
	"errors"
	"fmt"
	"gocv.io/x/gocv"
	"honor_money/cmd"
	"honor_money/glade"
	"honor_money/tool"
	"time"
)
//这个文件使用各种标志图，判断当前可能存在的异常情况，比如网络断开，好友邀请游戏等并尝试处理，并返回到首页重新开始流程
//CheckIndex 检测是否到达首页
func CheckIndex(dev *cmd.AdbHandler,img gocv.Mat)(ok bool,err error){
	//对比首页图片
	devImgs:=Imgs[dev.Size.String()]
	fmt.Println("匹配首页")

	ok,_,err=tool.MatchPicture(devImgs["index"],img)
	if err!=nil{
		glade.NoticeError("登录图片匹配错误:"+err.Error())
		return false,nil
	}
	if ok{

		CStep.Step=INDEX
		//执行到达对应任务页面的逻辑//不成功就不算做成功
		 return TaskLogic(dev)
	}else{
		return false,nil
	}
}
//检测当前是否有异常状态
func CheckMain(dev *cmd.AdbHandler,img gocv.Mat)(ok bool,err error){
	//对比首页图片
	//根据各种图片判断当前位置，暂时只做冒险，完成后添加六国远征逻辑
	//判断是否是链接丢失
	//读取图片
	devImgs:=Imgs[dev.Size.String()]

	ok,p,err:=tool.MatchPicture(devImgs["confirm"],img)
	if err!=nil{
		return false,errors.New("match confirm pic failed")
	}
	if ok{
		dev.Tap(p.X,p.Y)
		return true,nil
	}
	//读取红色关闭图片

	ok,p,err =tool.MatchPicture(devImgs["red_confirm"],img)
	if err!=nil{
		return false,errors.New("match redconfirm pic failed")
	}
	if ok{
		dev.Tap(p.X,p.Y)
		return true,nil
	}
	//点击继续
	fmt.Println("点击屏幕继续")
	ok,p,err =tool.MatchPicture(devImgs["continue"],img)
	if err!=nil{
		return false,errors.New("match continue pic failed")
	}
	if ok{
		dev.Tap(p.X,p.Y)
		return true,nil
	}
	fmt.Println("点击屏幕继续")
	//点击返回按钮

	ok,p,err =tool.MatchPicture(devImgs["returnbtn"],img)
	if err!=nil{
		return false,errors.New("match returnbtn pic failed")
	}
	if ok{
		dev.Tap(p.X,p.Y)
		return true,nil
	}
	//判断是否需要返回

	hasReturn:=false //判断是否匹配到过return
	End:
	for{

		ok,p,err:=tool.MatchPicture(devImgs["return"],img)
		if err!=nil{
			return false,errors.New("match return pic failed")
		}
		if ok{
			hasReturn=true
			dev.Tap(p.X,p.Y)
		}else {
			if hasReturn{
				return ok,nil
			}else{
				break End
			}
		}
		img ,err= glade.PullPicture(dev)
		if err!=nil{

			return false,err
		}
	}
	//判断是否需要点叉叉

	hasClose:=false //判断是否匹配到过return
End1:
	for{


		ok,p,err:=tool.MatchPicture(devImgs["close"],img)
		if err!=nil{
			return false,errors.New("match close pic failed")
		}
		if ok{
			hasClose=true
			dev.Tap(p.X+5,p.Y+5)
		}else {
			if hasClose{
				return ok,nil
			}else{
				break End1
			}
		}
		img ,err= glade.PullPicture(dev)
		if err!=nil{

			return false,err
		}
	}
	//不常见的黄色返回
	//点击返回按钮

	ok,p,err =tool.MatchPicture(devImgs["yreturn"],img)
	if err!=nil{
		return false,errors.New("match yreturn pic failed")
	}
	if ok{
		dev.Tap(p.X,p.Y)
		return true,nil
	}
	return false,nil
}

//TaskLogic 首页到冒险模式逻辑
func TaskLogic(dev *cmd.AdbHandler)(ok bool,err error){
	devImgs:=Imgs[dev.Size.String()]
	dev.ConfTap("index_wan")//点击万象天工
	time.Sleep(1*time.Second)
	//滑动到万象天工选项卡
	dev.ConfSwipe("wan_swipe")//滑动关卡
	//查询万象天工标识
	time.Sleep(4*time.Second)

	//多次尝试拉取图片,除非图片不为空
	taskImg,err :=glade.PullPicture(dev)
	if err!=nil{

	    return  false,err
	}
	if taskImg.Empty(){
		return false, errors.New("读取手机图片失败")
	}
	defer func(taskImg gocv.Mat) {
		fmt.Println("关闭taskimg1")
		taskImg.Close()
		fmt.Println("关闭taskimg2")
	}(taskImg)

	//获取 万象天工特征图
	ok,p,err := tool.MatchPicture(devImgs["risk"],taskImg)
	if !ok{
		return false,nil
	}
	dev.Tap(p.X,p.Y)
	//开始冒险模式流程
	time.Sleep(2*time.Second)
	dev.ConfTap("to_risk")
	time.Sleep(2*time.Second)
	dev.Step = 0
	StartTaskRist(dev)

	return true,nil
}

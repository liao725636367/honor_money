package initialize

import (
	"fmt"
	"honor_money/cmd"
	"honor_money/def"
	"honor_money/glade"
	"honor_money/tool"
	"strconv"
	"time"
)

//开始冒险模式流程
func StartTaskRist(dev *cmd.AdbHandler){
	//开启主定时器，由主定时器决定是否需要开启图片识别
	if dev.Timers[0] == nil {
			fmt.Println("主定时器1")
		dev.Timers[0] = tool.UserTicker(4000, MainHandler, dev)
		fmt.Println("主定时器1")
	}

	fmt.Println(dev.Timers)

	def.Entry1.SetSensitive(false)
}


/**
执行主要进程
*/
func MainHandler(dev *cmd.AdbHandler) {
	fmt.Println("当前step", dev.Step)
	switch dev.Step {
	case 0: //第一步开始闯关
		Step0(dev)
	case 1:
		Step1(dev)
	case 2:
		Step2(dev)
	case 3:
		Step3(dev)
	}
	//if dev.Timers[0].Flag{
	//	return true
	//}else{
	//	return false
	//}
}




/**
图片识别线程 修改步数
*/
func PicStepFix(dev *cmd.AdbHandler) {
	devImgs:=Imgs[dev.Size.String()]
	//var flag bool
	//if dev.Timers[1]!=nil&&dev.Timers[1].Flag{
	//	flag= true
	//}else{
	//	flag= false
	//}
	//判断是第四步才匹配
	if dev.Step != 2 {
		return
	}
	//上传图片
	//获取匹配图片
	img,err:=glade.PullPicture(dev)
	//获取匹配图片

	//获取区域
	areaX := 0
	areaY := 0
	if area, ok := dev.PointConf["rect_area1"]; ok {
		areaX = area[0]
		areaY = area[1]
	} else {
		glade.NoticeError("图片区域配置获取失败")
		fmt.Println("图片区域配置获取失败")
		//glade.MsgError(def.AppWin1,"图片区域配置获取失败")
		StopTask(dev,PAUSE)
		return
	}
	_ = areaX
	_ = areaY //暂时用不到此配置
	//跟已有图片匹配
	ok, point, err := tool.MatchPicture(devImgs["try"], img)
	if err != nil {
		glade.NoticeError("匹配图片失败:"+err.Error())
		fmt.Println(def.AppWin1, "匹配图片失败:"+err.Error())
		return
	}
	if ok {
		dev.ChangeStep(3)
	} else if time.Since(dev.StepTime) > 300*time.Second {
		fmt.Println("进行通用图片识别纠错")
		StopTask(dev,CHANGE)
		err = PointStep(dev) //如果此步骤时间太长就判断是否异常
		if err!=nil{
			fmt.Println("错误是:",err)
		    glade.NoticeError(err.Error())
		}
	}
	fmt.Printf("匹配坐标:%d,%d\n", point.X, point.Y)
	//成功就设置为第二步
	//否则继续
	return
}



//滑动关卡点击关卡
func Step0(dev *cmd.AdbHandler) {

	//滑动
	dev.ConfSwipe("swipe0")
	time.Sleep(time.Millisecond * 500)
	//点击三个地方
	//dev.ConfTap("tap1_3")
	//time.Sleep(time.Millisecond * 300)
	//dev.ConfTap("tap1_0")
	//time.Sleep(time.Millisecond * 300)
	//dev.ConfTap("tap1_1")
	//time.Sleep(time.Millisecond * 300)
	dev.ConfTap("tap1_2")
	time.Sleep(time.Millisecond * 300)
	dev.ChangeStep(1)
}

//点击开始闯关
func Step1(dev *cmd.AdbHandler) {
	dev.ConfTap("tap2")
	time.Sleep(time.Millisecond * 5) //这里时间久一点，在这里创建图片识别线程
	if dev.Timers[1] == nil {
		dev.Timers[1] = tool.UserTicker(6000, PicStepFix, dev)
	}
	dev.ChangeStep(2)
}

//点击对话框和再次挑战
func Step2(dev *cmd.AdbHandler) {
	dev.ConfTap("tap3")
}

//Step3 判断是否需要结束循环和其它结束后需要的操作
func Step3(dev *cmd.AdbHandler) {
	dev.Nums--
	//减少显示数量
	def.Entry1.SetText(strconv.Itoa(dev.Nums))
	if dev.Nums <= 0 {
		StopTask(dev,CHANGE)
	}
	dev.ChangeStep(1) //空操作
}
func ChangeStep(dev *cmd.AdbHandler, step int) {
	//修改时间这里做处理时间
}


package cmd

import (
	"errors"
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"regexp"
	"strconv"
	"strings"
	"time"
)
//使用两个定时器 一个用于图片识别 一个用于步数循环控制
type TimeHandler struct {
	Handle glib.SourceHandle
	Flag bool
}

//AdbHandler 封装每个设备为一个结构体，有操作该设备的方法实现
type AdbHandler struct {
	DevId string //设备id
	Nums int //刷新次数
	RealSize *Size //屏幕真实尺寸
	Size *Size //屏幕尺寸
	PointConf map[string][]int
	Step int //当前设备执行到的步数
	//Timers [2]* TimeHandler //定时器信息
	Timers [9]chan bool //定时器信息
	StepTime time.Time //记录每一步持续多久
	IsEnd bool
}

type Size struct {
	Width int
	Height int
}
func (s Size) String()string{
	return fmt.Sprintf("%dx%d",s.Width,s.Height)
}
//Init 初始化设备需要的信息
func (dev *AdbHandler)Init() (err error){
	_,err =dev.GetSize()
	if err!=nil{
		return err
	}

	dev.Step=0
	dev.StepTime=time.Now()
	dev.Timers=[9]chan bool{nil,nil,nil,nil,nil,nil,nil,nil,nil}
	dev.IsEnd=false
	//判断设备尺寸 因为每种设备分辨率不一样，都统一设置 1080x1920 方便适配
	if dev.RealSize.String()!="1080x1920"{
		dev.Cmd("shell wm size 1080x1920")
		//adb shell wm density 480 //修改dpi 这个可以不改
		dev.Size.Width=1080
		dev.Size.Height=1920
	}
	sizeStr:=fmt.Sprintf("%dx%d",dev.Size.Width,dev.Size.Height)
	if conf,ok:=PointerMap[sizeStr];ok{
		dev.PointConf= conf
		return nil
	}else{

		return errors.New("不支持的分辨率:"+sizeStr)
	}

}
//GetDevices 获取所有设备号，提供给用户选择设备
func GetDevices()(devs []string,err error){
	//执行adb命令
	res,err:=RunCmd("adb devices")
	if err!=nil{
	    fmt.Printf("get devs failed,err:%v",err)
	    return nil,err
	}
	//结果进行分割
	tmpDevs:=strings.Split(res,"\n")//临时存储切割字符串
	//正则表达式定义
	reg:=regexp.MustCompile(`^([a-zA-Z0-9\.:-]{5,})`)
	//正则匹配出设备号
	for _,str:=range tmpDevs{
		str=strings.TrimSpace(str)//去除多余的空格

		str1 :=reg.FindString(str)
		//fmt.Printf("%#v **** %#v *** %#v\n",str,str1,len(str1))

		if len(str1) >=  5{
			devs=append(devs,str1)
		    //fmt.Printf("failed,err:%v",err)

		}
	}
	//返回设备号列表
	return devs,nil
}
//Shell 执行当前设备操作命令
func (dev *AdbHandler) Cmd(cmd string)(ok bool,err error){
	_,err =RunCmd(fmt.Sprintf(" adb -s %s  %s ",dev.DevId,cmd))
	if err!=nil{ //这里设计的是失败就返回错误
	    fmt.Printf("failed,err:%v res:%s",err)
	    return false,err
	}else{
		return true,nil
	}
}
//ShellInput 执行设备input系列命令
func (dev *AdbHandler) ShellInput(shell string)(ok bool,err error){

	return dev.Cmd(fmt.Sprintf("shell input %s",shell))//封装参数直接传给更基本的Shell方法
}
//Tap 执行设备点击命令
func (dev *AdbHandler) Tap(x,y int)(ok bool,err error){

	return dev.ShellInput(fmt.Sprintf(" tap %d %d",x,y))//封装参数直接传给更基本的Shell方法
}
//Tap 执行设备滑动命令ax,ay 起点坐标 bx by 结束点坐标 dur 延时毫秒
func (dev *AdbHandler) Swipe(ax,ay,bx,by,dur int)(ok bool,err error){

	return dev.ShellInput(fmt.Sprintf(" swipe %d %d %d %d %d ",ax,ay,bx,by,dur))//封装参数直接传给更基本的Shell方法
}
//ConfTap 根据配置点击设备
func (dev *AdbHandler) ConfTap(str string)(ok bool,err error){
	 if point,ok:=dev.PointConf[str];ok{
		return dev.Tap(point[0],point[1])
	 }else{
	 	return false,errors.New("config: "+str+" not found")
	 }
}
//ConfSwipe 根据配置滑动设备
func (dev *AdbHandler) ConfSwipe(str string)(ok bool,err error){
	if point,ok:=dev.PointConf[str];ok{
		return dev.Swipe(point[0],point[1],point[2],point[3],point[4])
	}else{
		return false,errors.New("config: "+str+" not found")
	}
}

//GetSize 获取设备尺寸信息
func (dev *AdbHandler) GetSize()(size *Size,err error){
	if dev.Size!=nil{
		return dev.Size,nil
	}
	res,err:=RunCmd(fmt.Sprintf(" adb -s %s shell  wm size ",dev.DevId))
	if err!=nil{
		fmt.Println("cmd 错误",err)
	    return
	}
	reg:=regexp.MustCompile("(Physical size:)(.*)([0-9]+)(x)([0-9]+)")
	strSplit:=reg.FindStringSubmatch(res)
	//if len(strSplit) !=2{
	//	err =errors.New("get device size failed")
	//	return
	//}
	//因为手机游戏都是以长边为宽
	width,err:=strconv.Atoi(strSplit[3])
	if err!=nil{
	    return
	}

	height,err:=strconv.Atoi(strSplit[5])
	if err!=nil{
	    return
	}
	if width <1 ||height <1{
		err=errors.New(fmt.Sprintf("宽高:%dx%d不合法",width,height))
	}
	dev.RealSize=&Size{
		Width: width,
		Height: height,
	}
	dev.Size=&Size{
		Width: width,
		Height: height,
	}


	return dev.Size,nil

}

func (dev *AdbHandler)ChangeStep(step int){
	dev.Step =step
	dev.StepTime=time.Now()
}

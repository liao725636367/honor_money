package initialize

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"honor_money/cmd"
	"honor_money/def"
	"honor_money/glade"
	"strconv"
)

func InitEvent(){
	def.Button1.Connect("clicked", func(btn *gtk.Button) {
		btn1Click(btn)
		//只能在回调里调用窗口这里打印
	})

	def.Button2.Connect("clicked", func(btn *gtk.Button) {
		btn2Click(btn)

	})
	def.ComboBoxText1.Connect("changed", func(combox *gtk.ComboBoxText) {
		comboxChange(combox)
	})
	def.Entry1.Connect("changed", func(entry *gtk.Entry) {
		entry1Change(entry)
	})
}
func entry1Change(entry *gtk.Entry){
	conStr,err  := entry.GetText()
	if err!=nil{
		conStr  ="0"
	}
	//对输入信息控制
	con,err:=strconv.Atoi(conStr)
	if err!=nil{
	    con=0
	}
	if con > 9999999{
		con = 9999999
	}
	conStr =strconv.Itoa(con)

	entry.SetText(conStr)

}
func btn1Click(btn *gtk.Button){
	//暂时只做单设备，后面可以多设备 这里定义一下 后面做多设备好扩充
	dev:=def.AdbDevice
	win:=def.AppWin1
	label,err:=btn.GetLabel()
	if err!=nil{
		glade.MsgError(win,err.Error())
	    return
	}

	if label=="开始"{
		if def.AdbDevice.DevId == ""{
			glade.MsgError(win,"请选择设备")
			return
		}
		//获取设置挑战次数
		numS,err:=def.Entry1.GetText()
		if err!=nil{
			glade.MsgError(win,err.Error())
		    return
		}
		num,err:=strconv.Atoi(numS)
		if err!=nil{
			num =0
		}
		def.AdbDevice.Nums =num
		if def.AdbDevice.Nums < 1{
			glade.MsgError(win,"次数已完成")
			return
		}
		//初始化设备
		 err=def.AdbDevice.Init()
		if err!=nil{
			glade.MsgError(win,"初始化设备失败"+err.Error())
		    return
		}
		fmt.Println("开始",def.AdbDevice.Nums)
		  StartTask(dev)
		 def.AppWin1.SetTitle("金币助手-"+dev.DevId)
		def.Button1.SetLabel("暂停")

	}else{
		fmt.Println("暂停")
		StopTask(dev,PAUSE)
		def.Button1.SetLabel("开始")
	}

}
//btn2Click 获取设备按钮点击事件
func btn2Click(btn *gtk.Button){
	fmt.Println("获取设备")
	appWin:=def.AppWin1
	comboBoxText:=def.ComboBoxText1
	devs,err:=cmd.GetDevices()
	if err!=nil{
		glade.MsgError(appWin, err.Error())
		return
	}else if len(devs) <1{
		glade.MsgError(appWin, "设备为空")
		return
	}
	//获取之前选择的设备index
	actIndex:=comboBoxText.GetActiveID()//获取之前选项
	fmt.Printf("actIndex:%#v\n",actIndex)
	comboBoxText.RemoveAll()//移除所有
	comboBoxText.Append("0",def.NoneSelect)
	for index:=range devs{
		curIndex:=strconv.Itoa(100+index)
		//如果没有选择设备就选择当前设备
		if actIndex=="0"||actIndex==""{
			actIndex=curIndex
		}
		fmt.Printf("curIndex:%#v\n",curIndex)
		comboBoxText.Append(curIndex,devs[index] )
	}
		comboBoxText.SetActiveID(actIndex)

}
//comboxChange 设备选择改变后修改当前设备
func comboxChange(combox *gtk.ComboBoxText){
	did:=combox.GetActiveText()
	if did==def.NoneSelect{
		did=""
	}
	dev:=&cmd.AdbHandler{
		DevId: did,
		Nums:  0,
		Size:  nil,
		RealSize:nil,
	}//设置设备id
	if did!=""&&did!=def.AdbDevice.DevId{
		//这里不初始化设备因为用户还没决定开始刷
		def.AdbDevice=dev
		fmt.Println("当前设备号:",def.AdbDevice.DevId,"设备尺寸:",def.AdbDevice.Size)
	}



}

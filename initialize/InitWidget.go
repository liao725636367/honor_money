package initialize

import (
	"honor_money/def"
	"honor_money/glade"
)

//先初始化window才能初始化组件
func InitWidget()(err error){
	//按钮事件监听
	builder:=def.Builder1
	def.Button1,err =glade.GetButton(builder,"button1")
	if err!=nil{
		return err
	}
	def.Button2,err =glade.GetButton(builder,"button2")
	if err!=nil{
	    return err
	}
	//获取下拉框
	def.ComboBoxText1,err =glade.GetComboBoxText(builder,"comboboxtext1")
	if err!=nil{
	    return err
	}
	comboBoxText:=def.ComboBoxText1
	comboBoxText.SetWrapWidth(100)
	comboBoxText.SetActive(0)
	//次数输入框
	def.Entry1,err=glade.GetEntry(builder,"entry1")
	if err!=nil{
		return  err
	}
	//刷完关机按钮
	def.Check1,err=glade.GetCheckButton(builder,"checkbutton1")
	if err!=nil{
	    return  err
	}
	//提示栏
	def.Notice,err=glade.GetLabel(builder,"label3")
	if err!=nil{
		return  err
	}
	//事件监听
	InitEvent()
	return nil
}


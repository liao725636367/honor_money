package glade

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"honor_money/def"
)
//GetFixed 获取固定布局容器
func GetFixed(builder *gtk.Builder,name string)(*gtk.Fixed,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	fixed,ok:=buttonObj.(*gtk.Fixed)
	if !ok{
		return nil,err
	}
	return fixed,nil
}
//GetEntry 获取输入框
func GetEntry(builder *gtk.Builder,name string)(*gtk.Entry,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	entry,ok:=buttonObj.(*gtk.Entry)
	if !ok{
		return nil,err
	}
	return entry,nil
}
//GetEntry 选择框
func GetCheckButton(builder *gtk.Builder,name string)(*gtk.CheckButton,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	checkButton,ok:=buttonObj.(*gtk.CheckButton)
	if !ok{
		return nil,err
	}
	return checkButton,nil
}
//GetButton 获取按钮
func GetButton(builder *gtk.Builder,name string)(*gtk.Button,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	button,ok:=buttonObj.(*gtk.Button)
	if !ok{
		return nil,err
	}
	return button,nil
}
//GetComboBoxText 获取文本下拉框
func GetComboBoxText(builder *gtk.Builder,name string)(*gtk.ComboBoxText,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	comboBoxText,ok:=buttonObj.(*gtk.ComboBoxText)
	if !ok{
		return nil,err
	}
	return comboBoxText,nil
}
//GetLabel 获取文字展示框
func GetLabel(builder *gtk.Builder,name string)(*gtk.Label,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	label,ok:=buttonObj.(*gtk.Label)
	if !ok{
		return nil,err
	}
	return label,nil
}
//成功弹窗
func MsgSuccess(appWin *gtk.Window, content string){
	dialog:=gtk.MessageDialogNew(appWin,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		content)
	dialog.SetTitle("成功")
	//dialog.SetIcon(def.SuccessIcon)
	 dialog.Run()
	//if flag ==gtk.RESPONSE_YES{
	//	fmt.Println("你准备好了")
	//}else if flag == gtk.RESPONSE_NO{
	//	fmt.Println("你还没准备好")
	//}else{
	//	fmt.Println("你什么都没选择，还关闭了我")
	//}
	defer dialog.Destroy()
}
//失败弹窗
func MsgError(appWin *gtk.Window, content string){
	dialog:=gtk.MessageDialogNew(appWin,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_ERROR,
		gtk.BUTTONS_OK,
		content)
	dialog.SetTitle("错误")
	//使用success图标

	//dialog.SetIcon(def.ErrorIcon)
	dialog.Run()
	//if flag ==gtk.RESPONSE_YES{
	//	fmt.Println("你准备好了")
	//}else if flag == gtk.RESPONSE_NO{
	//	fmt.Println("你还没准备好")
	//}else{
	//	fmt.Println("你什么都没选择，还关闭了我")
	//}
	defer dialog.Destroy()
}

//显示错误信息
func NoticeError(con string){
	mark:=fmt.Sprintf("<span  weight='bold' font_desc='12' foreground='#FF5722'>%s</span> ",con)
	def.Notice.SetMarkup(mark)
}
//显示错误信息
func NoticeSuccess(con string){
	mark:=fmt.Sprintf("<span weight='bold'  font_desc='12' foreground='#5FB878 '>%s</span>",con)
	def.Notice.SetMarkup(mark)
}


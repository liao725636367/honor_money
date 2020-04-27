package glade

import (
	"github.com/gotk3/gotk3/gtk"
)

/**
获取应用窗体
*/
func GetApplicationWindow(builder *gtk.Builder,name string)(*gtk.ApplicationWindow,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	//fmt.Println(reflect.TypeOf(buttonObj))
	ApplicationWindow,ok:=buttonObj.(*gtk.ApplicationWindow)
	if !ok{
		return nil,err
	}
	return ApplicationWindow,nil
}


/**
获取普通窗体
*/
func GetWindow(builder *gtk.Builder,name string)(*gtk.Window,error){
	buttonObj,err:=builder.GetObject(name)
	if err!=nil{
		return nil,err
	}
	//fmt.Println(reflect.TypeOf(buttonObj))
	window,ok:=buttonObj.(*gtk.Window)
	if !ok{
		return nil,err
	}
	return window,nil
}


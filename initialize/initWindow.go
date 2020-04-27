package initialize

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"honor_money/data"
	"honor_money/def"
	"honor_money/glade"
)

//InitWindow 窗口初始化操作
func InitWindow()(err error){
	def.Builder1,err=glade.GetBuilderFromString(data.GetTpl())
	fmt.Println("没错误1")
	if err!=nil{
	    return err
	}
	builder:=def.Builder1
	//事件链接
	builder.ConnectSignals(nil)
	fmt.Println("没错误2")
	def.AppWin1,err =glade.GetWindow(builder,"window1")
	if err!=nil{
	    return err
	}
	fmt.Println("没错误3")
	appWin:=def.AppWin1
	fmt.Println("没错误4")
	//_=buf
	//设置居中
	appWin.SetPosition(gtk.WIN_POS_CENTER)
	appWin.SetTitle("金币助手")
	appWin.SetResizable(false)//禁止全屏和调整窗口大小
	appWin.Connect("destroy", func() { //窗口关闭处理
		gtk.MainQuit()
	})
	return nil
}




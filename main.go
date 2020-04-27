package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"honor_money/def"
	"honor_money/initialize"
	"log"
)

func main() {
		onactive( )
}
//onactive 程序启动处理
func onactive( ){
	gtk.Init(nil)
	//初始化窗口属性信息
	err:=initialize.InitWindow()
	errorCheck(err)
	//初始化组件信息
	err=initialize.InitWidget()
	errorCheck(err)
	//展示窗口并阻塞等待事件
	def.AppWin1.ShowAll()
	//窗口等待
	gtk.Main()
	//绑定窗口事件
	//执行循环判断游戏进度逻辑
}
func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		fmt.Println("出现异常")
		log.Panic(e)
	}
}

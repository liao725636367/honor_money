package def
//全局变量定义
import (
	"github.com/gotk3/gotk3/gtk"
	"honor_money/cmd"
)

var Button1 *gtk.Button //开始停止按钮
var Button2 *gtk.Button //获取设备按钮
var Entry1 *gtk.Entry //刷次数输入框
var Check1 *gtk.CheckButton //刷完关机
var ComboBoxText1 *gtk.ComboBoxText //下拉选择按钮
var Notice *gtk.Label//通知区域
var AppWin1 *gtk.Window //窗口
var Builder1 *gtk.Builder //glade 模板对象
var NoneSelect string  ="--请选择--"//空设备默认选择
var AdbDevice = &cmd.AdbHandler{DevId:"",Nums:0} //初始化为空设备
var GamePackName ="com.tencent.tmgp.sgame"//游戏包名
var GamePackFullName="com.tencent.tmgp.sgame/com.tencent.tmgp.sgame.SGameActivity"//游戏包全名



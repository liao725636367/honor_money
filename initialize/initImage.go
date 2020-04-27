package initialize

import (
	"fmt"
	"gocv.io/x/gocv"
	"honor_money/cmd"
)
//将需要匹配图片资源初始化
var Imgs = make(map[string]map[string]gocv.Mat)
func InitImage(dev *cmd.AdbHandler) (err error) {
	s := dev.Size.String()
	if _, ok := Imgs[s]; ok {
		return nil
	}
	devImgs := make(map[string]gocv.Mat)
	imgPath := fmt.Sprintf("./image/%s/", s)
	index := gocv.IMRead(imgPath+"index.png", gocv.IMReadColor)
	if index.Empty() {
		return err
	}
	devImgs["index"]=index
	closeImg := gocv.IMRead(imgPath+"close.png", gocv.IMReadColor)
	if closeImg.Empty() {
		return err
	}
	devImgs["close"]=closeImg

	redConfirm := gocv.IMRead(imgPath+"red_confirm.png", gocv.IMReadColor)
	if redConfirm.Empty() {
		return err
	}
	devImgs["red_confirm"]=redConfirm

	ret := gocv.IMRead(imgPath+"return.png", gocv.IMReadColor)
	if ret.Empty() {
		return err
	}
	devImgs["return"]=ret

	risk:= gocv.IMRead(imgPath+"risk.png", gocv.IMReadColor)
	if risk.Empty() {
		return err
	}
	devImgs["risk"]=risk

	confirm := gocv.IMRead(imgPath+"confirm.png", gocv.IMReadColor)
	if confirm.Empty() {
		return err
	}
	devImgs["confirm"]=confirm

	try := gocv.IMRead(imgPath+"try.png", gocv.IMReadColor)
	if try.Empty() {
		return err
	}
	devImgs["try"]=try

	login := gocv.IMRead(imgPath+"login.png", gocv.IMReadColor)
	if login.Empty() {
		return err
	}
	devImgs["login"]=login
	cont := gocv.IMRead(imgPath+"continue.png", gocv.IMReadColor)
	if cont.Empty() {
		return err
	}
	devImgs["continue"]=cont
	returnbtn := gocv.IMRead(imgPath+"returnbtn.png", gocv.IMReadColor)
	if returnbtn.Empty() {
		return err
	}
	devImgs["returnbtn"]=returnbtn
	yreturn := gocv.IMRead(imgPath+"yreturn.png", gocv.IMReadColor)
	if yreturn.Empty() {
		return err
	}
	devImgs["yreturn"]=yreturn
	Imgs[s]=devImgs


	return nil
}

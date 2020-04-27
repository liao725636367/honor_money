package glade

import (
	"errors"
	"fmt"
	"gocv.io/x/gocv"
	"honor_money/cmd"
	"path/filepath"
	"regexp"
)

func PullPicture(dev *cmd.AdbHandler)(mat gocv.Mat,err error){
	//设备名称过滤 替换掉 非单词字符
	devName := regexp.MustCompile(`\W`).ReplaceAll([]byte(dev.DevId), []byte(""))
	//获取设备图片
	mat = gocv.NewMat()
	picName := fmt.Sprintf("sc-%s.png", devName)
	_, err  = dev.Cmd(fmt.Sprintf("shell screencap -p /sdcard/%s && adb -s %s  pull /sdcard/%s ", picName, dev.DevId, picName))
	if err != nil {
		fmt.Println("上传设备图片失败:" + err.Error())
		//glade.MsgError(def.AppWin1,"上传设备图片失败:"+err.Error())
		return mat,errors.New("上传设备图片失败:" + err.Error())
	}
	//读取图片
	picPath := filepath.Join("./", picName)
	img := gocv.IMRead(picPath, gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("读取图片失败")
		//glade.MsgError(def.AppWin1,)
		return mat,errors.New("读取图片失败" )
	}
	return  img,nil
}

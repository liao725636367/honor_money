package tool

import "C"
import (
	"errors"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"reflect"
	"time"
)

//UserTicker 自定义定时器，使用channel实现控制定时器的关闭
func UserTicker(dur int,callback interface{}, args ...interface{})chan bool{
	ticker := time.NewTicker(time.Duration(dur)*time.Millisecond)
	fun := reflect.ValueOf(callback)
	if fun.Kind() != reflect.Func {
		panic("not a function")
	}
	vargs:=make([]reflect.Value,len(args))
	for i,arg:=range args{
		vargs[i]=reflect.ValueOf(arg)
	}

	stopChan := make(chan bool,2)//这里多给一个方便函数本身调取chan
	go func(ticker *time.Ticker,fun reflect.Value, vargs []reflect.Value ) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fun.Call(vargs)
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(ticker,fun,vargs)

	return stopChan
}
//MatchPicture 图片特征匹配是否匹配
func MatchPicture(tempPic gocv.Mat, destPic gocv.Mat) (ok bool, p *image.Point, err error) {
	var point = &image.Point{X: 0, Y: 0}

	if tempPic.Rows() > destPic.Rows() || tempPic.Cols() > destPic.Cols() {
		return false, point, errors.New("模板图不能比对比图大")
	}

	if tempPic.Empty() || destPic.Empty() {
		return false, point, errors.New("模板图或目标图为空")
	}

	var result = gocv.NewMat()
	var mask = gocv.NewMat()
	var method gocv.TemplateMatchMode
	defer func() { //完成后回收此函数产生资源
		result.Close()
		mask.Close()

	}()
	//标准相关匹配 抗干扰效果最好
	method = gocv.TmCcoeffNormed
	gocv.MatchTemplate(tempPic, destPic, &result, method, mask)
	minVal, maxVal, minLoc, maxLoc := gocv.MinMaxLoc(result) //计算查找结果
	//标准化差值平方和匹配 方式结果是 值越小越匹配，需要区分开
	if method == gocv.TmSqdiffNormed {
		fmt.Println("相似度", minVal)
		if minVal < 0.1 {
			point.X = minLoc.X
			point.Y = minLoc.Y
			return true, point, nil
		} else {
			return false, point, nil
		}
	} else {
		fmt.Println("相似度", maxVal)
		if maxVal > 0.9 {//0.9差不多算相似
			//获取匹配的坐标，进行点击操作
			point.X = maxLoc.X
			point.Y = maxLoc.Y
			return true, point, nil
		} else {
			return false, point, nil
		}
	}
}

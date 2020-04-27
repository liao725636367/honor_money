package glade

import (
	"github.com/gotk3/gotk3/gtk"
)

/**
从文件中获取获取builder对象
*/
func GetBuilder(filename string)(*gtk.Builder,error){
	builder,err:=gtk.BuilderNew()
	if err !=nil{
		return nil,err
	}
	err = builder.AddFromFile(filename)
	if err!=nil{
		return nil,err
	}
	return builder,nil
}
func GetBuilderFromString(str string)(*gtk.Builder,error){
	builder,err:=gtk.BuilderNew()
	if err !=nil{
		return nil,err
	}
	err = builder.AddFromString(str)
	if err!=nil{
		return nil,err
	}
	return builder,nil
}
//PixBUfFromAssetData 从静态资源里面获取gtk文件
//func PixBUfFromAssetData(filename string)(pixBuf *gdk.Pixbuf,err error){
//	Data,err :=data.Asset(filename)
//	fmt.Println("没错2-1")
//	if err!=nil{
//		return nil,err
//	}
//	fmt.Println("没错2-2")
//	pixBufLoader,err :=gdk.PixbufLoaderNew()
//	if err!=nil{
//		return nil,err
//	}
//	fmt.Println("没错2-3")
//	defer pixBufLoader.Close()
//	pixBuf,err =pixBufLoader.WriteAndReturnPixbuf(Data)
//	if err!=nil{
//		return nil,err
//	}
//	fmt.Println("没错2-4")
//	return  pixBuf,nil
//	//return nil,nil
//
//}

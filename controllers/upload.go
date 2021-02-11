package controllers

import (
	"fmt"
	"github.com/astaxie/goredis"
	"io"
	"log"
	"os"
	"path/filepath"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Post() {
	fileData, fileHeader, err := c.GetFile("upload")

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message":"上传失败"} //没有得到文件的信息
		c.ServeJSON()
		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
	fmt.Println(fileData)
	//now := time.Now()
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileTypeFlag := 0
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中,好像差不多这些格式，多了再加就行
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" ||  fileExt == ".jpeg" {
		fileTypeFlag = 1
	}
	//文件路径
	fileDir := "static\\img"
	fileName := fileHeader.Filename
	filePathStr := filepath.Join(fileDir, fileName)
	fmt.Println("Path:",filePathStr)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0,"message":"上传失败"}
		c.ServeJSON()
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0,"message":"上传失败"}
		c.ServeJSON()
		return
	}
	if fileTypeFlag == 1 {
		//存入数据库中
		var client goredis.Client
		client.Addr = "127.0.0.1:6379"
		err := client.Set("Pic",[]byte(filePathStr))
		if err != nil {
			log.Println(err)
		}
	}
	c.Data["json"] = map[string]interface{}{"code": 1,"message":"上传成功"}
	c.ServeJSON()
}
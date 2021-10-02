package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
)

func UploadFile(c *gin.Context,fileName string) (string,error) {
	file,err := c.FormFile(fileName)
	if err != nil {
		return "",err
	}

	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":true,
		".png":true,
		".gif":true,
		".jpeg":true,
	}
	if _,ok := allowExtMap[extName];!ok {
		return "",errors.New("文件名后缀不合法")
	}
	// 创建图片保存目录
	day := GetDay()
	dir := "./static/back/upload" + day

	err = os.MkdirAll(dir,0666)
	if err != nil {
		return "",err
	}

	// 生成文件名称和文件保存目录
	name := strconv.FormatInt(GetUnix(),10) + extName
	dst := path.Join(dir,name)

	err = c.SaveUploadedFile(file,dst)

	return dst,nil
}


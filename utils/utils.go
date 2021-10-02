package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// 时间戳转换成日期

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h,str)
	return fmt.Sprintf("%x",h.Sum(nil))
}

//获取时间戳

func GetUnix() int64 {
	return time.Now().Unix()
}

// 字符串转int

func Int(str string) (int,error) {
	n,err := strconv.Atoi(str)
	return n,err
}

//获取当前的日期

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

//获取年月日

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}


// int转换成string

func String(n int) string  {
	str := strconv.Itoa(n)
	return str
}



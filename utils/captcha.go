package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)


var store = base64Captcha.DefaultMemStore

func Captcha()(id string, b64s string,err error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height: 40,
		Width: 100,
		NoiseCount: 0,
		ShowLineOptions: 2 | 4,
		Length: 4,
		Source: "sfjdhghgfdfkhsfdksfjh",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver,store)
	id,b64s,err = c.Generate()
	return id,b64s,err
}
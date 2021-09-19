package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

func Captcha()(id string,) {
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
		}
	}
}
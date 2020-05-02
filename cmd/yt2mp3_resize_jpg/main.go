package main

import (
	"fmt"
	"os"
	"github.com/carstenkuckuk/yt2mp3/lib"
)

func main(){
	nArgs := len(os.Args)
	if (nArgs<3) || (nArgs>5) {
		fmt.Println("Usage: yt2mp3_resize_jpg <Source-Jpg-Filename> <Target-Jpeg-Filename> <Number_of_Horizontal_Pixels_Default_300> <Number_of_Vertical_Pixels_Default_300>")
	} else {
		sourceJpgFile := os.Args[1]
		targetJpgFile := os.Args[2]
		xPixels := "300"
		yPixels := "300"
		if (nArgs>3){
			xPixels = os.Args[3]
		}
		if (nArgs>4){
			yPixels = os.Args[4]
		}
		yt2mp3.ResizeJPEGFileToSize(sourceJpgFile, targetJpgFile, xPixels, yPixels)
	}

}

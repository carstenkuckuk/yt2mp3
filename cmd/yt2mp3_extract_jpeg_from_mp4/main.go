package main

import (
	"fmt"
	"os"
	"github.com/carstenkuckuk/yt2mp3/lib"
)

func main(){
	if len(os.Args) != 4 {
		fmt.Println("Usage: yt2mp3_extract_jpeg_from_mp4 <MP4-Filename> <Timecode> <JPEG-Filename>")
	} else {
		mp4filename := os.Args[1]
		timecode := os.Args[2]
		jpegfilename := os.Args[3]
		yt2mp3.ExtractJpegFromMP4(mp4filename, timecode, jpegfilename)
	}

}

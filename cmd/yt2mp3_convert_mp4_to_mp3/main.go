package main

import (
	"fmt"
	"os"
	"github.com/carstenkuckuk/yt2mp3/lib"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("Usage: yt2mp3_convert_mp4_to_mp3 <MP4-Filename> <MP3-Filename")
	} else {
		mp4filename := os.Args[1]
		mp3filename := os.Args[2]
		yt2mp3.ConvertMP4FileToMP3File(mp4filename, mp3filename)
	}

}

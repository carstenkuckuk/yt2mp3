package main

import (
	"fmt"
	"os"
	"github.com/carstenkuckuk/yt2mp3/lib"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("Usage: yt2mp3_download_as_mp4 <Youtube-URL> <MP4-Filename>")
	} else {
		url := os.Args[1]
		mp4filename := os.Args[2]
		yt2mp3.DownloadVideoFromYoutubeURLToMp4File(url, mp4filename)
	}

}

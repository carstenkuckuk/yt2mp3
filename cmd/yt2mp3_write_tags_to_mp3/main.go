package main

import (
	"fmt"
	"os"
	"github.com/carstenkuckuk/yt2mp3/lib"
)

func main(){
	if len(os.Args) != 7 {
		fmt.Println("Usage: yt2mp3_write_tags_to_mp3 <MP3-Filename> <Album> <Artist> <Title> <Download-URL-As-Comment> <JPEG-Filename>")
	} else {
		mp3filename := os.Args[1]
		album := os.Args[2]
		artist := os.Args[3]
		title := os.Args[4]
		url := os.Args[5]
		jpegfilename := os.Args[6]
		yt2mp3.WriteID3V2TagsToMP3File(mp3filename, album, artist, title, url, jpegfilename)
	}

}

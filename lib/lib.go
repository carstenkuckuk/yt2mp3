package yt2mp3

import (
	"fmt"
	"github.com/bogem/id3v2"
	. "github.com/kkdai/youtube"
	"os/exec"
)

func DownloadVideoFromYoutubeURLToMp4File(YoutubeUrl string, MP4Filename string){
	fmt.Println("DownloadVideoFromYoutubeURLToMp4File: URL=", YoutubeUrl, ", Filename=", MP4Filename)

	// Benutze kkdai/youtube um eine MP4-Datei von Youtube herunterzuladen:
	y := NewYoutube(true)
	err := y.DecodeURL(YoutubeUrl)
	fmt.Println("err:", err)
	y.StartDownload( MP4Filename );
	fmt.Println("err:", err)
}

func ConvertMP4FileToMP3File(MP4Filename string, MP3Filename string){
	// Benutze externes Tool C:\bin\ffmpeg zum Konvertieren von mp4 nach mp3:
	fmt.Println("Vor dem Exec")
	cmd := exec.Command("c:\\bin\\ffmpeg.exe", "-i", MP4Filename, MP3Filename)
	myerr:=cmd.Run()
	fmt.Println("Nach dem exec", myerr)
}

func WriteID3V2TagsToMP3File(MP3Filename string , Album string , Artist string , Title string, Comment string){
	// Schreibe id3v2-Tags:
	tag, _ := id3v2.Open(MP3Filename, id3v2.Options{Parse:true})
	defer tag.Close()
	tag.SetAlbum(Album)
	tag.SetArtist(Artist)
	tag.SetTitle(Title)
	comment := id3v2.CommentFrame{
		Encoding:   id3v2.EncodingUTF8,
		Language:    "eng",
		Description: "Source",
		Text:        Comment,
	}
	tag.AddCommentFrame(comment)
	tag.Save()
}

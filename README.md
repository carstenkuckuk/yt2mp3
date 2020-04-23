# yt2mp3
Tools to download videos from Youtube, turn them into MP3 files, and embed ID3V2 information. This is a
little project I created while I started learning the go language. The heavy lifting is done by the
referenced projects kkdai/youtube, bogem/id3v2, and ffmpeg. My code is just glue code.

## Compiling
First, install the GO SDK. Then clone this repository here, and build the executables:
```
git clone https://github.com/carstenkuckuk/yt2mp3.git
cd yt2mp3
go build ./...
```

## Usage
Make sure that you have an ffmpg executable in your path. Then you can use all the commands that are
available under ./cmd

## Download a video from Youtube to an mp4 files
```
yt2mp3_download_as_mp4.exe https://www.youtube.com/watch?v=pXRux1bwnFI Louane_Midi_sur_novembre_ft_Julien_Dore.mp4
```

## Extract the audio track of an MP4 file and write it to an MP3 file
```
yt2mp3_convert_mp4_to_mp3.exe Louane_Midi_sur_novembre_ft_Julien_Dore.mp4 Louane_Midi_sur_novembre_ft_Julien_Dore.mp3 
```



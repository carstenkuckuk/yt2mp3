package main

import (
	"fmt"
	. "github.com/carstenkuckuk/yt2mp3/lib"
)


func main(){
	fmt.Println("Hello World!")
	
	DownloadVideoFromYoutubeURLToMp4File("https://www.youtube.com/watch?v=LYAvhujK4nA", "Dalida_AlainDeloin_Paroles_paroles.mp4")
	ExtractJpeg300x300FromMP4("Dalida_AlainDeloin_Paroles_paroles.mp4", "00:00:10", "Dalida_AlainDeloin_Paroles_paroles.jpg")
	ConvertMP4FileToMP3File("Dalida_AlainDeloin_Paroles_paroles.mp4","Dalida_AlainDeloin_Paroles_paroles.mp3")
	WriteID3V2TagsToMP3File("Dalida_AlainDeloin_Paroles_paroles.mp3", "YoutubeDownloadsAlbum", "Dalida & Alain Deloin", "Paroles, paroles", "https://www.youtube.com/watch?v=LYAvhujK4nA", "Dalida_AlainDeloin_Paroles_paroles.jpg" )

	DownloadVideoFromYoutubeURLToMp4File("https://www.youtube.com/watch?v=X3FF_rGNI50", "FaadaFreddy_Deezer_LeRing.mp4")
	DownloadVideoFromYoutubeURLToMp4File("https://www.youtube.com/watch?v=xVfqm1fQ-nU", "ChristopheObispo_LesMotsBleus.mp4")
	DownloadVideoFromYoutubeURLToMp4File("https://www.youtube.com/watch?v=UXnJzrMKBK8", "AgnesObel_DeezerSession.mp4")
	DownloadVideoFromYoutubeURLToMp4File("https://www.youtube.com/watch?v=oUGrWlfz9Ww", "LEJ_LeRingLive.mp4")
	DownloadVideoFromYoutubeURLToMp4File("https://www.youtube.com/watch?v=pXRux1bwnFI", "Louane_Midi_sur_novembre_ft_Julien_Dore.mp4")
	
	ExtractJpeg300x300FromMP4("FaadaFreddy_Deezer_LeRing.mp4", "00:02:22", "FaadaFreddy_Deezer_LeRing.jpg")
	ExtractJpeg300x300FromMP4("ChristopheObispo_LesMotsBleus.mp4", "00:01:30", "ChristopheObispo_LesMotsBleus.jpg")
	ExtractJpeg300x300FromMP4("AgnesObel_DeezerSession.mp4", "00:01:20", "AgnesObel_DeezerSession.jpg")
	ExtractJpeg300x300FromMP4("LEJ_LeRingLive.mp4", "00:01:35", "LEJ_LeRingLive.jpg")
	ExtractJpeg300x300FromMP4("Louane_Midi_sur_novembre_ft_Julien_Dore.mp4", "00:00:31", "Louane_Midi_sur_novembre_ft_Julien_Dore.jpg")
	
	ConvertMP4FileToMP3File("FaadaFreddy_Deezer_LeRing.mp4","FaadaFreddy_Deezer_LeRing.mp3")
	ConvertMP4FileToMP3File("ChristopheObispo_LesMotsBleus.mp4", "ChristopheObispo_LesMotsBleus.mp3")
	ConvertMP4FileToMP3File("AgnesObel_DeezerSession.mp4", "AgnesObel_DeezerSession.mp3" )
	ConvertMP4FileToMP3File("LEJ_LeRingLive.mp4", "LEJ_LeRingLive.mp3")
	ConvertMP4FileToMP3File("Louane_Midi_sur_novembre_ft_Julien_Dore.mp4", "Louane_Midi_sur_novembre_ft_Julien_Dore.mp3")

	WriteID3V2TagsToMP3File("FaadaFreddy_Deezer_LeRing.mp3", "YoutubeDownloadsAlbum", "Faada Freddy", "Deezer Le Ring", "https://www.youtube.com/watch?v=X3FF_rGNI50", "FaadaFreddy_Deezer_LeRing.jpg" )
	WriteID3V2TagsToMP3File("ChristopheObispo_LesMotsBleus.mp3", "YoutubeDownloadsAlbum", "Christophe & Pascal Obispo", "Les mots bleus", "https://www.youtube.com/watch?v=xVfqm1fQ-nU", "ChristopheObispo_LesMotsBleus.jpg" )
	WriteID3V2TagsToMP3File("AgnesObel_DeezerSession.mp3", "YoutubeDownloadsAlbum", "Agnes Obel", "Deezer Session", "https://www.youtube.com/watch?v=UXnJzrMKBK8", "AgnesObel_DeezerSession.jpg")
	WriteID3V2TagsToMP3File("LEJ_LeRingLive.mp3", "YoutubeDownloadsAlbum", "L.E.J", "Deezer Le Ring Live", "https://www.youtube.com/watch?v=oUGrWlfz9Ww", "LEJ_LeRingLive.jpg")
	WriteID3V2TagsToMP3File("Louane_Midi_sur_novembre_ft_Julien_Dore.mp3", "YoutubeDownloadsAlbum", "Louane", "Midi sur novembre ft. Julien Doré", "https://www.youtube.com/watch?v=pXRux1bwnFI", "Louane_Midi_sur_novembre_ft_Julien_Dore.jpg")

}

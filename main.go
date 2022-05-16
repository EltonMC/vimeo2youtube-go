package main

import (
	"github.com/joho/godotenv"
	"log"
	"main/lib"
	"os"
	"strconv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	vimeoVideoId, _ := strconv.Atoi(os.Getenv("VIMEO_VIDEO_ID"))
	video := lib.GetVideo(os.Getenv("VIMEO_TOKEN"), vimeoVideoId)
	filename := lib.DownloadVideo(lib.GetVideoDownloadLink(*video))
	lib.UploadVideoYouTube(filename)

}

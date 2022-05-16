package lib

import (
	"github.com/cavaliergopher/grab/v3"
	"github.com/silentsokolov/go-vimeo/vimeo"
	"golang.org/x/oauth2"
	"log"
)

func GetVideo(token string, id int) *vimeo.Video {
	video, _, err := getClientVimeo(token).Videos.Get(id)
	if err != nil {
		log.Fatalf("Error on vimeo api: %v", err)
	}
	return video
}

func GetVideoDownloadLink(video vimeo.Video) string {
	downloadList := video.Download
	bestQuality := downloadList[len(downloadList)-1]

	return bestQuality.Link
}

func DownloadVideo(url string) string {
	resp, err := grab.Get("./download/", url)
	if err != nil {
		log.Fatalf("Error on download video %v", err)
	}

	return resp.Filename
}

func getClientVimeo(token string) *vimeo.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return vimeo.NewClient(tc, nil)
}

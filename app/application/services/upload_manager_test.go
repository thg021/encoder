package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/thg021/encoder/application/services"
)

func init() {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatalf("Error loading env variables")
	}
}

func TestVideoServiceUpload(t *testing.T) {
	video, repo := Prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("fcprojetovideo")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = "fcprojetovideo"
	videoUpload.VideoPath = os.Getenv("localStoragePath") + "/" + video.ID

	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(50, doneUpload)

	result := <-doneUpload

	require.Equal(t, result, "upload complete")

	err = videoService.Finish()
	require.Nil(t, err)

}

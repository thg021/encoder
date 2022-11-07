package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/thg021/encoder/application/repositories"
	"github.com/thg021/encoder/application/services"
	"github.com/thg021/encoder/domain"
	"github.com/thg021/encoder/framework/database"
)

func Prepare() (*domain.Video, repositories.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "boas_vindas.mp4"
	video.ResourceID = "resource_id"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepositoryDb(db)
	return video, repo
}

func init() {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatalf("Error loading env variables")
	}
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := Prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("fcprojetovideo")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

}

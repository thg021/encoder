package repositories_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/thg021/encoder/application/repositories"
	"github.com/thg021/encoder/domain"
	"github.com/thg021/encoder/framework/database"
)

func TestNewVideoRepository(t *testing.T) {
	db := database.NewDbTest()

	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.ResourceID = "resource_id"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepositoryDb(db)
	repo.Insert(video)

	v, err := repo.Find(video.ID)
	require.Nil(t, err)
	require.NotEmpty(t, v.ID)
	require.Equal(t, v.ID, video.ID)
}

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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.ResourceID = "resource_id"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepositoryDb(db)
	repo.Insert(video)

	job, err := domain.NewJob("output", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepositoryDb(db)
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, j.ID)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryUpdate(t *testing.T) {
	db := database.NewDbTest()

	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.ResourceID = "resource_id"
	video.CreatedAt = time.Now()
	repo := repositories.NewVideoRepositoryDb(db)
	repo.Insert(video)
	job, err := domain.NewJob("output", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepositoryDb(db)
	job.Status = "Completed"
	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	require.Nil(t, err)
	require.NotEmpty(t, j.ID)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Status, "Completed")
}

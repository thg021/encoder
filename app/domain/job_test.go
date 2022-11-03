package domain_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/thg021/encoder/domain"
)

func TestNewJob(t *testing.T) {

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "id"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("output", "status", video)

	require.Nil(t, err)
	require.NotNil(t, job)
	require.Equal(t, "output", job.OutputBucketPath)
	require.Equal(t, "status", job.Status)

}

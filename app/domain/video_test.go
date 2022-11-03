package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thg021/encoder/domain"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	//esta e uma conversao de quem trabalha com GO
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIDisNotUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "is_not_valid_id"
	video.ResourceID = "video"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

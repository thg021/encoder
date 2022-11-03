package domain

import "time"

//este e o responsavel por fazer o processamento do nosso video
type Job struct {
	ID               string
	OutputBucketPath string
	Status           string
	Video            *Video
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

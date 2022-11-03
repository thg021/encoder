package domain

import "time"

type Video struct {
	ID         string // este e nosso id, interno do sistema
	ResourceID string //este id é do sistema que enviara as informação
	FilePath   string
	CreatedAt  time.Time
}

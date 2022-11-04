package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"` // este e nosso id, interno do sistema /// O nome do video convertido ser o nome da pasta na nuvem
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`           //este id é do sistema que enviara as informação
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"-" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func init() {
	//roda primeiro
	govalidator.SetFieldsRequiredByDefault(true) ///se por padrao os campos da struct for requerida ele vai validar
}

func NewVideo() *Video {
	return &Video{}
}

func (v *Video) Validate() error {
	_, err := govalidator.ValidateStruct(v)

	if err != nil {
		return err
	}

	return nil
}

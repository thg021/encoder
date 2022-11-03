package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `valid:"uuid"`    // este e nosso id, interno do sistema
	ResourceID string    `valid:"notnull"` //este id é do sistema que enviara as informação
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
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

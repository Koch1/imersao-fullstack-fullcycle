package model

import (
	"github.com/asaskevich/govalidator" 
	 uuid "gihub.com/satori/go.uuid"
	 "time" 
)
func init(){
	govalidator.SetFieldsRequiredByDefault(value: true);
}

type Base struct {
	ID strint `json:"id" valid:"uuid"`
	CreateadAt time.Time `json:"created_at" valid:"-"`
	UpdateaAt time.Time `json:"updated_at" valid:"-"`

}
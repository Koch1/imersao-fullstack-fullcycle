package model

import (
	   "github.com/asaskevich/govalidator" 
		uuid "gihub.com/satori/go.uuid"
		"time" 
)

type User struct {
	Base `value:"required"`
	Name string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"email"`
}
func (user *User) isValid() error{
	validator, err:= govalidator.ValidateStruct(user)	
	if err!=nil{
		return err
	}
	return nil
}


func NewUser (name string,email string)(*User, error){
	user :=User{
		Name:name,
		Email:email,
	}
	user.ID= uuid.NewV4().String()
	user.CreateadAt=time.Now()
	err:=user.isValid()
	if err!=nil{
		return nil, err
	}
	return &user,nil 
}
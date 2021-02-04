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
	Account []*Account `valid:"-"`	
}
func (user *User) isValid() error{
	_ , err:= govalidator.ValidateStruct(bank)

	
	if err!=nil{
		return err
	}
	return nil
}


func NewBank (code string,name string)(*Bank, error){
	bank :=Bank{
		Code:code,
		Name:name,
	}
	bank.ID= uuid.NewV4().String()
	bank.CreateadAt=time.Now()
	err:=bank.isValid()
	if err!=nil{
		return nil, err
	}
	return &bank,nil 
}
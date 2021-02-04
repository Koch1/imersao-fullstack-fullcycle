package model

import (
	   "github.com/asaskevich/govalidator" 
		uuid "gihub.com/satori/go.uuid"
		"time" 
)

type Bank struct {
	Base `value:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json:"name"`
	Account []*Account `valid:"-"`	
}
func (bank *Bank) isValid() error{
	validator, err:= govalidator.ValidateStruct(bank)
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
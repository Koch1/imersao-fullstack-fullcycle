package model

import (
	"github.com/asaskevich/govalidator" 
	 uuid "gihub.com/satori/go.uuid"
	 "time" 
)
//Criacao do banco de dados
type Account struct{
	Base `valid:"required"`
	DwnerName string `json:"owner_name" valid:"notnull"`
	Bank *Bank `valid:"-"`
	Number string `json:"number" valid:"notnull"`
	PixKey []*PixKey `valid:"-"`
}

func(account *Account) isValid() error{
	_,err := govalidator.ValidateStruct(account)
	if err!=nil{
		return err
	}
	return nil
}

//Iniciar uma entidade
func NewAccount(bank *Bank, number string, ownerName string)(*Account , error) {
	account := Account {
		OwnerName: ownerName,
		Bank: bank,
		Number: number,
	}
	account.ID= uuid.NewV4().String()
	account.CreateadAt=time.Now()
	err:= account.isValid()
	if err != nil {
		return nil, err
	}
	return &account,nil

}
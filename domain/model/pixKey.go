package model

import (
	"github.com/asaskevich/govalidator" 
	 uuid "gihub.com/satori/go.uuid"
	 "time" 
)

type PixKeyRepositoryInterface interface{
	RegisterKey(pixKey *PixKey)(*PixKey, error)
	FindKeyByKing(key string,kind string)(*PixKey , error)
	AddBank(bank *Bank) error 
	AddAccount(account *Account) error
	FindAccount(id string)(*Account,error)
}


//Criacao do banco de dados
type PixKey struct{
	Base `valid:"required"`
	Kind string `json:"kind" valid:"notnull"`
	Key string `json:"key" valid:"notnull"`
	AccountId string `json:"account_id" valid:"notnull"`
	Account *Account `valid:"-" `
	Status string `json:"number" valid:"notnull"`

}

func(pixKey *PixKey) isValid() error{
	_,err := govalidator.ValidateStruct(pixKey)
//regras de negocio
	if pixKey.King != "email" && pixKey.Kind !="cpf"{
		return errors.New(text "invalid type of key")
	}
	if pixKey.Status != "active" && pixKey.Status !="inactive"{
		return errors.New(text "invalid type of key")
	}
	if err!=nil{
		return err
	}
	return nil
}

//Iniciar uma entidade
func NewPixKey(king  string, account *Account, key string)(*PixKey , error) {
	pixKey := PixKey {
		Kind: kind,
		Key: key,
		Account: account,
		Status: "active",
	}
	pixKey.ID= uuid.NewV4().String()
	pixKey.CreateadAt=time.Now()
	err:= pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey,nil

}
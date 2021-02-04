package model

import (
	   "github.com/asaskevich/govalidator" 
		uuid "gihub.com/satori/go.uuid"
		"time" 
)

const (
	TransactionPending string ="pending"
	TransactionCompleted string ="completed"
	TransactionError string ="error"
	TransactionConfirmed string ="confirmed"

)
type Transaction struct {
	Transaction []Transaction
}
type TransactionRepositoryInterface interface{
	Register (transaction *Transaction) errors
	Save(transaction *Transaction) error
	Find (id string) (*Transaction, error)

}

func (transaction *Transaction) Complete() error{
	transaction.Status=TransactionCompleted
	transaction.UpdateaAt=time.Now()
	err:=t.isValid() 
	return err
}
func (transaction *Transaction) Confirm() error{
	transaction.Status=TransactionConfirmed
	transaction.UpdateaAt=time.Now()
	err:=t.isValid() 
	return err
}
func (transaction *Transaction) Cancel( description String) error{
	transaction.Status=TransactionError
	transaction.UpdateaAt=time.Now()
	transaction.Description=description
	err:=t.isValid() 
	return err
}

type Transaction struct {
	Base `value:"required"`
	AccountFrom *Account `valid:"-"`
	Amount float64 `json:"amount" valid:"notnull"`
	PixKeyTo *PixKey `valid:"-"`
	Status string `json:"status" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	CancelDescription string `json:"cancel_description"`
}
func (transaction *Transaction) isValid() error{
	_, err:= govalidator.ValidateStruct(transaction)
	if transaction.Amount <= 0{
		return errors.New(text "the amount must be greater thant 0")
	}
	if transaction.Status!=TransactionPending && transaction.Status!=TransactionCompleted && transaction.Status!=TransactionError{
		return errors.New(text "invalid status for the transaction")
	}
	if transaction.PixKeyTo.AccountId==transaction.AccountFrom.ID{
		return errors.New(text "invalid status for the transaction")
	}
	if err!=nil{
		return err
	}
	return nil
}


func NewTransaction (accountFrom *Account, amount float64,pixKeyTo *PixKey,description string)(*Bank, error){
	transaction :=Transaction{
		Base `value:"required"`
		AccountFrom:  accountFrom,
		Amount: amount,
		PixKeyTo: pixKeyTo,
		Status: TransactionPending
		Description: description,
	}
	transaction.ID= uuid.NewV4().String()
	transaction.CreateadAt=time.Now()
	err:=transaction.isValid()
	if err!=nil{
		return nil, err
	}
	return &transaction,nil 
}
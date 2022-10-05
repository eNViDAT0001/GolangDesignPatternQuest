package creational

import "errors"

type account struct {
	username string
	balance  int64
	ID       uint64
}

type Account interface {
	WithDraw(amount int64) error
	Deposit(amount int64) error
	GetBalance() (int64, error)
	GetInfo() account
}
type AccountBuilder interface {
	WithUserName(username string) AccountBuilder
	WithBalance(balance int64) AccountBuilder
	WithID(id uint64) AccountBuilder
	Build() Account
}

func (a *account) WithDraw(amount int64) error {
	if amount > a.balance {
		return errors.New("amount is not enough")
	}
	a.balance -= amount
	return nil
}
func (a *account) Deposit(amount int64) error {
	a.balance += amount
	return nil
}
func (a *account) GetBalance() (int64, error) {
	return a.balance, nil
}
func (a *account) GetInfo() account {
	return *a
}
func (a *account) WithUserName(username string) AccountBuilder {
	a.username = username
	return a
}
func (a *account) WithBalance(balance int64) AccountBuilder {
	a.balance = balance
	return a
}

func (a *account) WithID(id uint64) AccountBuilder {
	a.ID = id
	return a
}
func (a *account) Build() Account {
	return a
}
func NewAccountBuilder() AccountBuilder {
	return &account{}
}

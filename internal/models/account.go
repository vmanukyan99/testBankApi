package models

import (
	"errors"
	"gorm.io/gorm"
	"math"
)

type Account struct {
	gorm.Model
	ID         uint64
	Sum        float64
	CurrencyId int
	Currency   Currency `gorm:"foreignKey:CurrencyId"`
}

func (a *Account) AddFunds(sum float64) error {
	if sum < 0 {
		return errors.New("the amount cannot be less than zero")
	}
	a.Sum += sum
	return updateAccount(a)
}

func (a *Account) SumProfit() error {
	a.Sum *= 1.06
	return updateAccount(a)
}

func (a *Account) Withdraw(sum float64) error {
	if sum <= 0.7*a.Sum {
		a.Sum -= sum
		return updateAccount(a)
	}
	return errors.New("amount is too large to withdraw. Please reduce the amount and try again")
}

func (a *Account) GetCurrency() (string, error) {
	if a.Currency.Name == "" {
		return a.Currency.Name, errors.New("error getting Account Currency")
	}
	return a.Currency.Name, nil
}

func (a Account) GetAccountCurrencyRate(currency string) (float64, error) {
	var currencyInfo Currency
	DataBase.First(&currencyInfo, "name = ?", currency)
	if currencyInfo.Name == "" {
		return 0, errors.New("currency not found")
	}
	return currencyInfo.Rate / a.Currency.Rate, nil
}

func (a Account) GetBalance(currency string) (float64, error) {
	var currencyInfo Currency
	DataBase.First(&currencyInfo, "name = ?", currency)
	if currencyInfo.Rate == 0 {
		return 0, errors.New("currency not found")
	}
	return math.Round(a.Sum * currencyInfo.Rate * 100) / 100, nil
}

func AccountSelection(accountId string) (*Account, error) {
	var result *Account
	DataBase.First(&result, accountId)
	if result.ID == 0 {
		return result, errors.New("account not found")
	}
	return result, nil
}

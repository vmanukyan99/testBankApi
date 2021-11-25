package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DataBase = connect()

func connect() *gorm.DB {
	if _, err := os.Stat("bank.sqlite"); err == nil {
		dataBase, openError := gorm.Open(sqlite.Open("bank.sqlite"), &gorm.Config{})
		if openError != nil {
			panic(openError.Error())
		}
		return dataBase
	}
	return create()
}

func create() *gorm.DB {
	_, createError := os.Create("bank.sqlite")
	if createError != nil {
		panic(createError.Error())
	}

	dataBase, openError := gorm.Open(sqlite.Open("bank.sqlite"), &gorm.Config{})
	if openError != nil {
		panic(openError.Error())
	}

	createTableCurrenciesError := dataBase.Migrator().CreateTable(&Currency{})
	if createTableCurrenciesError != nil {
		createTableCurrenciesError.Error()
	}

	createTableAccountsError := dataBase.Migrator().CreateTable(&Account{})
	if createTableAccountsError != nil {
		createTableAccountsError.Error()
	}

	dataBase.Create(&Currency{Name: "SBP", Rate: 1})
	dataBase.Create(&Currency{Name: "RUB", Rate: 0.7523})
	dataBase.Create(&Account{Sum: 0, CurrencyId: 1})

	return dataBase
}

func updateAccount(a *Account) error {
	DataBase.Model(&a).Update("sum", a.Sum)
	DataBase.Model(&a).Update("currency_id", a.CurrencyId)
	return DataBase.Error
}

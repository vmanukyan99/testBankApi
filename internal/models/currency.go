package models

type Currency struct {
	ID   uint64
	Name string
	Rate float64
}

var SBP Currency
var RUB Currency

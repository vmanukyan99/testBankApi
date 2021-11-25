package handlers

import (
	"net/http"
	"testBankApi/internal"
	"testBankApi/internal/models"
)

func GetCurrency(w http.ResponseWriter, r *http.Request) {
	currentAccount, accountError := models.AccountSelection(r.URL.Query()["account"][0])
	if accountError != nil {
		internal.SendResponse(accountError.Error(), w)
		return
	}

	currency, getCurrencyError := currentAccount.GetCurrency()
	if getCurrencyError != nil {
		internal.SendResponse(getCurrencyError.Error(), w)
		return
	}

	internal.SendResponse(currency, w)
}

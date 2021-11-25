package handlers

import (
	"net/http"
	"testBankApi/internal"
	"testBankApi/internal/models"
)

func GetAccountCurrencyRate(w http.ResponseWriter, r *http.Request) {
	currentAccount, accountError := models.AccountSelection(r.URL.Query()["account"][0])
	if accountError != nil {
		internal.SendResponse(accountError.Error(), w)
		return
	}

	cur := "SBP"
	if r.URL.Query()["currency"] != nil {
		cur = r.URL.Query()["currency"][0]
	}

	rate, getAccountCurrencyError := currentAccount.GetAccountCurrencyRate(cur)
	if getAccountCurrencyError != nil {
		internal.SendResponse(getAccountCurrencyError.Error(), w)
		return
	}

	internal.SendResponse(rate, w)
}

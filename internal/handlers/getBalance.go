package handlers

import (
	"net/http"
	"testBankApi/internal"
	"testBankApi/internal/models"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	currentAccount, accountError := models.AccountSelection(r.URL.Query()["account"][0])
	if accountError != nil {
		internal.SendResponse(accountError.Error(), w)
		return
	}

	cur := "SBP"
	if r.URL.Query()["currency"] != nil {
		cur = r.URL.Query()["currency"][0]
	}

	balance, gettingError := currentAccount.GetBalance(cur)
	if gettingError != nil {
		internal.SendResponse(gettingError.Error(), w)
		return
	}

	internal.SendResponse(balance, w)
}

package handlers

import (
	"net/http"
	"strconv"
	"testBankApi/internal"
	"testBankApi/internal/models"
)

func Withdraw(w http.ResponseWriter, r *http.Request) {
	currentAccount, accountError := models.AccountSelection(r.URL.Query()["account"][0])
	if accountError != nil {
		internal.SendResponse(accountError.Error(), w)
		return
	}

	sum, parseError := strconv.ParseFloat(r.URL.Query()["sum"][0], 64)
	if parseError != nil {
		internal.SendResponse(parseError.Error(), w)
		return
	}

	withdrawError := currentAccount.Withdraw(sum)
	if withdrawError != nil {
		internal.SendResponse(withdrawError.Error(), w)
		return
	}

	internal.SendResponse("withdrawal was successful", w)
}

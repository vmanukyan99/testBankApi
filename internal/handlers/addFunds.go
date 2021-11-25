package handlers

import (
	"net/http"
	"strconv"
	"testBankApi/internal"
	"testBankApi/internal/models"
)

func AddFunds(w http.ResponseWriter, r *http.Request) {
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

	addFundsError := currentAccount.AddFunds(sum)
	if addFundsError != nil {
		internal.SendResponse(addFundsError.Error(), w)
		return
	}

	profitError := currentAccount.SumProfit()
	if profitError != nil {
		internal.SendResponse(profitError.Error(), w)
		return
	}

	internal.SendResponse("account has been successfully replenished", w)
}

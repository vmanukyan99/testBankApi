package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testBankApi/internal/handlers"
	"testBankApi/internal/models"
)

func main() {
	models.DataBase.First(&models.SBP, 1)
	models.DataBase.First(&models.RUB, 2)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	go http.HandleFunc("/rest/add_funds/", handlers.AddFunds)
	go http.HandleFunc("/rest/withdraw/", handlers.Withdraw)
	go http.HandleFunc("/rest/get_currency/", handlers.GetCurrency)
	go http.HandleFunc("/rest/get_account_currency_rate/", handlers.GetAccountCurrencyRate)
	go http.HandleFunc("/rest/get_balance/", handlers.GetBalance)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

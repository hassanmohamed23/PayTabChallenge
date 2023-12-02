package main

import (
	"log"
	"net/http"

	AccountController "github.com/hassanmohamed23/PayTabChallenge/internal/adapters/api"
	"github.com/hassanmohamed23/PayTabChallenge/internal/services"
	accountService "github.com/hassanmohamed23/PayTabChallenge/internal/services"
)

func main() {

	accounts, err := accountService.LoadAccounts("../accounts.json")
	if err != nil {
		log.Fatal(err)
	}

	api := AccountController.AccountController(services.TransferService{Accounts: accounts})
	log.Fatal(http.ListenAndServe(":8080", api.Router))
}

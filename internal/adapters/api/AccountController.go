package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/hassanmohamed23/PayTabChallenge/internal/models"
	accountService "github.com/hassanmohamed23/PayTabChallenge/internal/services"
)

type API struct {
	Router  *mux.Router
	Service *accountService.TransferService
}

func AccountController(Service accountService.TransferService) *API {

	api := &API{
		Router:  mux.NewRouter(),
		Service: &Service,
	}

	api.Router.HandleFunc("/accounts", api.listAccounts).Methods("GET")
	api.Router.HandleFunc("/transfer", api.transfer).Methods("POST")

	return api
}

// retrieve all accounts
func (api *API) listAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := api.Service.GetAccounts()
	respondWithJSON(w, http.StatusOK, accounts)
}

// money transfer endPoint
func (api *API) transfer(w http.ResponseWriter, r *http.Request) {
	var transferRequest model.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&transferRequest)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	response := api.Service.MakeTransfer(transferRequest)

	respondWithJSON(w, response.Code, response)
}

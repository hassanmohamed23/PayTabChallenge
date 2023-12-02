package services

import (
	"encoding/json"
	"net/http"
	"os"

	model "github.com/hassanmohamed23/PayTabChallenge/internal/models"
)

type TransferService struct {
	Accounts []model.Account
}

func NewTransferService(Accounts []model.Account) *TransferService {
	return &TransferService{
		Accounts: Accounts,
	}
}

func (s *TransferService) InitAccounts(accounts []model.Account) {

	s.Accounts = append(s.Accounts, accounts...)
}

func (s *TransferService) GetAccounts() []model.Account {

	if s.Accounts == nil || len(s.Accounts) == 0 {
		return []model.Account{}
	}

	return s.Accounts
}

func (s *TransferService) MakeTransfer(request model.TransferRequest) model.TransferResponse {

	senderIndex := findAccountIndexByAccountId(s.Accounts, request.From)
	beneficiaryIndex := findAccountIndexByAccountId(s.Accounts, request.To)

	if senderIndex == -1 || beneficiaryIndex == -1 {
		return model.TransferResponse{
			Code:    http.StatusBadRequest,
			Message: "Account not found.",
		}
	}

	senderAccount := &s.Accounts[senderIndex]
	beneficiaryAccount := &s.Accounts[beneficiaryIndex]

	if senderAccount.Balance < request.Amount {
		return model.TransferResponse{
			Code:    http.StatusBadRequest,
			Message: "insufficient funds",
		}
	}

	senderAccount.Balance -= request.Amount
	beneficiaryAccount.Balance += request.Amount

	return model.TransferResponse{
		Code:    http.StatusOK,
		Message: "Transfer successful",
	}
}

func findAccountIndexByAccountId(accounts []model.Account, accountId string) int {
	for i, acc := range accounts {
		if acc.ID == accountId {
			return i
		}
	}
	return -1
}

func LoadAccounts(filename string) ([]model.Account, error) {
	file, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	var accounts []model.Account
	err = json.Unmarshal(file, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/stellar/go/clients/horizonclient"
)

type Response struct {
	Balances []BalanceData `json:"balances"`
}

type BalanceData struct {
	AssetType string `json:"assetType"`
	TokenSymbol string `json:"tokenSymbol"`
	Balance string `json:"balance"`
}

func GetAccountBalances(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["accountID"]

	client := horizonclient.DefaultTestNetClient
	accountRequest := horizonclient.AccountRequest{AccountID: accountID}
	Account, err := client.AccountDetail(accountRequest)
	if err != nil {
		fmt.Println("Error loading account:", err)
		return
	}
	var balances []BalanceData
	for _, balance := range Account.Balances {
		// fmt.Printf("Asset type: %s, Token Symbol: %s, Balance: %s\n", balance.Asset.Type, balance.Code, balance.Balance)
		balances = append(balances, BalanceData{balance.Asset.Type, balance.Code, balance.Balance})
	}
	balancesResponse := Response{Balances: balances}

	jsonItems, err := json.Marshal(balancesResponse)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonItems)
}
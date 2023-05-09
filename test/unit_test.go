package test

import (
	"math/rand"
	"testing"
	"net/http"
	"net/http/httptest"
	"stellar_cli/controller"
	"stellar_cli/service"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

var (
	client = horizonclient.DefaultTestNetClient
	generalExpectation bool = true

	tokenSymbol string

	tokenTransferAmount string

	issuerPublicKey string = "GBTVCZ7WVG36YSBEF4X3QMJI6SBZ65JU5YAEXOCIZUSL6OAWUYORTD4Y"
	issuerSecret string = "SDJK34IP7MPIYLY7IAZ7LSNK2AJZY3CYGPIZIYK2LJWECL5RAASQL4ZZ"
	IssuerKeypair *keypair.Full = keypair.MustParseFull(issuerSecret)
	
	distributorPublicKey string 
	distributorSecret string
	DistributorKeypair *keypair.Full
	
	userPublicKey string = "GC3J24BHVNCV6OQIFOEDC6IPYQQ6MJZFP5XPFAGZN6L3T3C4COOFBCMX"
	userSecret string = "SCNUELCHDINKKRW2NKOUYJLK4OFD7NE26MWQ3HYGPIXDAJ5JFOULRGEK"
	UserKeypair *keypair.Full= keypair.MustParseFull(userSecret)
)


var charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// n is the length of random string we want to generate
func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func TestCreateNewToken(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tokenSymbol = randStr(5)

	t.Log("Issuing new Token and Distributor ...")

	distributorKP, err := keypair.Random()
	distributorPublicKey = distributorKP.Address()
	distributorSecret = distributorKP.Seed()
	DistributorKeypair = keypair.MustParseFull(distributorSecret)
	if err != nil {
		t.Error("Error :", err)
	}

	TransferXLMResponse := <-service.InitializeDefaultDistributorAccount(IssuerKeypair, distributorKP.Address(), client)
	transferDetail, err := client.TransactionDetail(TransferXLMResponse)
	if transferDetail.Successful != generalExpectation {
		t.Errorf("Expected %v but got %v", generalExpectation, transferDetail.Successful)
	}

	t.Log("Transaction Hashes : ")
	t.Log("InitializeDefaultDistributorAccount Transaction Hash : ", TransferXLMResponse)

	ChangeTrustResponse := <-service.ChangeTrust(distributorKP, IssuerKeypair.Address(), tokenSymbol, "100000", client)
	issuingDetail, err := client.TransactionDetail(ChangeTrustResponse)
	if issuingDetail.Successful != generalExpectation {
		t.Errorf("Expected %v but got %v", generalExpectation, issuingDetail.Successful)
	}
	if err != nil {
		t.Error("Error :", err)
	}
	t.Log("Issuing New Token Transaction Hash : ", ChangeTrustResponse)

	t.Log("Distributor Account Detail : ")
	t.Log("Distributor Address / Public Key : ", distributorKP.Address())
	t.Log("Distributor Secret / Seed Key : ", distributorKP.Seed())
	
	t.Log("User Input Arguments : ")
	t.Log("Token Name : ", tokenSymbol)
	t.Log("Issuer Secret / Seed Key : ", issuerSecret)
}

func TestMintTokenToDistributor(t *testing.T){
	MintTokenResponse := <-service.MintToken(IssuerKeypair, DistributorKeypair.Address(), "10", tokenSymbol, client)
	mintTokenDetail, err := client.TransactionDetail(MintTokenResponse)
	if mintTokenDetail.Successful != generalExpectation {
		t.Errorf("Expected %v but got %v", generalExpectation, mintTokenDetail.Successful)
	}
	if err != nil {
		t.Error("Error :", err)
	}
	t.Log("Mint Token Transaction Hash : ", MintTokenResponse)
}

func TestChangeTrustlineOfUser(t *testing.T) {
	ChangeTrustResponse := <-service.ChangeTrust(UserKeypair, issuerPublicKey, tokenSymbol, "10000", client)
	changeTrustDetail, err := client.TransactionDetail(ChangeTrustResponse)
	if changeTrustDetail.Successful != generalExpectation {
		t.Errorf("Expected %v but got %v", generalExpectation, changeTrustDetail.Successful)
	}
	if err != nil {
		t.Error("Error :", err)
	}
	t.Log("Change Trustline Transaction Hash : ", ChangeTrustResponse)
}

func TestTransferTokenFromDistributorToUser(t *testing.T) {
	tokenTransferAmount = "5.0000000"
	TransferToken := <-service.TransferToken(DistributorKeypair, userPublicKey, issuerPublicKey, tokenTransferAmount, tokenSymbol, client)
	transferTokenDetail, err := client.TransactionDetail(TransferToken)
	if transferTokenDetail.Successful != generalExpectation {
		t.Errorf("Expected %v but got %v", generalExpectation, transferTokenDetail.Successful)
	}
	if err != nil {
		t.Error("Error :", err)
	}
	t.Log("Transfer Token Transaction Hash : ", TransferToken)
	t.Log("Token transfer amount : ", tokenSymbol)
}

func TestCheckUserBalance(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/accounts/{accountID}/balance", http.HandlerFunc(controller.GetAccountBalances))

	// Create a new request with the fake route URL
	req, err := http.NewRequest("GET", "/accounts/" + userPublicKey + "/balance", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the fake route with the request and recorder
	r.ServeHTTP(rr, req)	

	// Check the balance for given symbol
    var resp controller.Response
    if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
        t.Errorf("error unmarshaling response: %v", err)
    }
    var tokenBalance string
    for _, balance := range resp.Balances {
        if balance.TokenSymbol == tokenSymbol {
            tokenBalance = balance.Balance
            break
        }
    }
    if tokenBalance != tokenTransferAmount {
		t.Error("Balance not match")
		t.Errorf("Expected %v but got %v", tokenTransferAmount, tokenBalance)
	} 
}


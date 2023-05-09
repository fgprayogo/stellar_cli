package service

import (
	"log"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

//InitializeDefaultDistributorAccount creates Distributor account
func InitializeDefaultDistributorAccount(SenderKeypair *keypair.Full, ReceiverPublicKey string,
	client *horizonclient.Client) <-chan string {
	
	res := make(chan string)

	go func() {
		defer close(res)
		// Get information about the Distributor account
		accountRequest := horizonclient.AccountRequest{AccountID: SenderKeypair.Address()}
		Account, err := client.AccountDetail(accountRequest)
		if err != nil {
			log.Fatal(err)
		}

		// Initialize the new Token
		op := txnbuild.CreateAccount{
			Destination: ReceiverPublicKey,
			Amount: "5",
		}
		
		// Construct the transaction that will carry the operation
		tx, err := txnbuild.NewTransaction(
			txnbuild.TransactionParams{
				SourceAccount: &Account,
				IncrementSequenceNum: true,
				Operations:    []txnbuild.Operation{&op},
				BaseFee: txnbuild.MinBaseFee,
				Preconditions:  txnbuild.Preconditions{TimeBounds: txnbuild.NewTimeout(300)}, // Use a real timeout in production!
			},
		)
		// Sign the transaction, serialise it to XDR, and base 64 encode it
		tx, err = tx.Sign(network.TestNetworkPassphrase, SenderKeypair)
		if err != nil {
			log.Fatal(err)
		}
		
		// Convert to base64
		txe, err := tx.Base64()
		if err != nil {
			log.Fatal(err)
		}

		// Submit the transaction
		resp, err := client.SubmitTransactionXDR(txe)
		if err != nil {
			hError := err.(*horizonclient.Error)
			log.Fatal("Error submitting transaction:", hError)
		}
		res <- resp.Hash
	}()

	return res
}
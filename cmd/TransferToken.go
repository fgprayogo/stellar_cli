/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"stellar_cli/service"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

func init() {
	// Create the 'CreateToken' command
	transferTokenCmd := &cobra.Command{
		Use:   "TransferToken",
		Short: "Create a new Stellar token",
		Run:   transferToken,
	}

	// Add the flags for 'CreateToken' command
	transferTokenCmd.Flags().StringVarP(&senderSecret, "sender", "s", "", "The secret key of the sender account")
	transferTokenCmd.MarkFlagRequired("sender")
	transferTokenCmd.Flags().StringVarP(&receiverPublicKey, "receiver", "r", "", "The receiver of the new token")
	transferTokenCmd.MarkFlagRequired("receiver")
	transferTokenCmd.Flags().StringVarP(&tokenSymbol, "token", "t", "", "The secret key of the token account")
	transferTokenCmd.MarkFlagRequired("token")
	transferTokenCmd.Flags().StringVarP(&issuerPublicKey, "issuer", "i", "", "The secret key of the issuer account")
	transferTokenCmd.MarkFlagRequired("issuer")
	transferTokenCmd.Flags().StringVarP(&amount, "amount", "a", "", "The secret key of the amount account")
	transferTokenCmd.MarkFlagRequired("amount")

	// Add the 'CreateToken' command to the root command
	rootCmd.AddCommand(transferTokenCmd)
}

func transferToken(cmd *cobra.Command, args []string){
	client := horizonclient.DefaultTestNetClient

	SenderKeypair := keypair.MustParseFull(senderSecret)

	TransferToken := <-service.TransferToken(SenderKeypair, receiverPublicKey, issuerPublicKey, amount, tokenSymbol, client)
	log.Println("Transfer Token Transaction Hash : ", TransferToken)

}
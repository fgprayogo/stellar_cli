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
	changeTrustlineCmd := &cobra.Command{
		Use:   "ChangeTrustline",
		Short: "Create a new Stellar token",
		Run:   changeTrusline,
	}

	// Add the flags for 'CreateToken' command
	changeTrustlineCmd.Flags().StringVarP(&accountSecret, "account", "n", "", "The account of the new token")
	changeTrustlineCmd.MarkFlagRequired("account")
	changeTrustlineCmd.Flags().StringVarP(&tokenSymbol, "token", "t", "", "The secret key of the token account")
	changeTrustlineCmd.MarkFlagRequired("token")
	changeTrustlineCmd.Flags().StringVarP(&issuerPublicKey, "issuer", "i", "", "The secret key of the issuer account")
	changeTrustlineCmd.MarkFlagRequired("issuer")
	changeTrustlineCmd.Flags().StringVarP(&limit, "limit", "l", "", "The secret key of the limit account")
	changeTrustlineCmd.MarkFlagRequired("limit")

	// Add the 'CreateToken' command to the root command
	rootCmd.AddCommand(changeTrustlineCmd)
}

func changeTrusline(cmd *cobra.Command, args []string){
	client := horizonclient.DefaultTestNetClient

	AccountKeypair := keypair.MustParseFull(accountSecret)

	ChangeTrustResponse := <-service.ChangeTrust(AccountKeypair, issuerPublicKey, tokenSymbol, limit, client)
	log.Println("Change Trustline Transaction Hash : ", ChangeTrustResponse)

}
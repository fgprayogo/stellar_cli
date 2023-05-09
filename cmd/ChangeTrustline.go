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
	// Create the 'ChangeTrustline' command
	changeTrustlineCmd := &cobra.Command{
		Use:   "ChangeTrustline",
		Short: "Change trustline",
		Run:   changeTrusline,
	}

	// Add the flags for 'ChangeTrustline' command
	changeTrustlineCmd.Flags().StringVarP(&accountSecret, "account", "n", "", "The account secret key")
	changeTrustlineCmd.MarkFlagRequired("account")
	changeTrustlineCmd.Flags().StringVarP(&tokenSymbol, "token", "t", "", "The token symbol")
	changeTrustlineCmd.MarkFlagRequired("token")
	changeTrustlineCmd.Flags().StringVarP(&issuerPublicKey, "issuer", "i", "", "The public key of the issuer account")
	changeTrustlineCmd.MarkFlagRequired("issuer")
	changeTrustlineCmd.Flags().StringVarP(&limit, "limit", "l", "", "The limit of the trust")
	changeTrustlineCmd.MarkFlagRequired("limit")

	// Add the 'ChangeTrustline' command to the root command
	rootCmd.AddCommand(changeTrustlineCmd)
}

func changeTrusline(cmd *cobra.Command, args []string){
	client := horizonclient.DefaultTestNetClient

	AccountKeypair := keypair.MustParseFull(accountSecret)

	ChangeTrustResponse := <-service.ChangeTrust(AccountKeypair, issuerPublicKey, tokenSymbol, limit, client)
	log.Println("Change Trustline Transaction Hash : ", ChangeTrustResponse)

}
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
	mintTokenCmd := &cobra.Command{
		Use:   "MintToken",
		Short: "Create a new Stellar token",
		Run:   mintToken,
	}

	// Add the flags for 'CreateToken' command
	mintTokenCmd.Flags().StringVarP(&issuerSecret, "issuer", "i", "", "The secret key of the issuer account")
	mintTokenCmd.MarkFlagRequired("issuer")
	mintTokenCmd.Flags().StringVarP(&distributorSecret, "distributor", "d", "", "The distributor of the new token")
	mintTokenCmd.MarkFlagRequired("distributor")
	mintTokenCmd.Flags().StringVarP(&tokenSymbol, "token", "t", "", "The secret key of the token account")
	mintTokenCmd.MarkFlagRequired("token")
	mintTokenCmd.Flags().StringVarP(&amount, "amount", "a", "", "The secret key of the amount account")
	mintTokenCmd.MarkFlagRequired("amount")

	// Add the 'CreateToken' command to the root command
	rootCmd.AddCommand(mintTokenCmd)
}

func mintToken(cmd *cobra.Command, args []string){
	client := horizonclient.DefaultTestNetClient

	IssuerKeypair := keypair.MustParseFull(issuerSecret)
	DistributorKeypair := keypair.MustParseFull(distributorSecret)

	MintToken := <-service.MintToken(IssuerKeypair, DistributorKeypair.Address(), amount, tokenSymbol, client)
	log.Println("Mint Token Transaction Hash : ", MintToken)

}
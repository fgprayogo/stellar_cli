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
	// Create the 'MintToken' command
	mintTokenCmd := &cobra.Command{
		Use:   "MintToken",
		Short: "Mint token",
		Run:   mintToken,
	}

	// Add the flags for 'MintToken' command
	mintTokenCmd.Flags().StringVarP(&issuerSecret, "issuer", "i", "", "The secret key of the issuer account")
	mintTokenCmd.MarkFlagRequired("issuer")
	mintTokenCmd.Flags().StringVarP(&distributorSecret, "distributor", "d", "", "The secret key of the distributor")
	mintTokenCmd.MarkFlagRequired("distributor")
	mintTokenCmd.Flags().StringVarP(&tokenSymbol, "token", "t", "", "The account symbol")
	mintTokenCmd.MarkFlagRequired("token")
	mintTokenCmd.Flags().StringVarP(&amount, "amount", "a", "", "The amount")
	mintTokenCmd.MarkFlagRequired("amount")

	// Add the 'MintToken' command to the root command
	rootCmd.AddCommand(mintTokenCmd)
}

func mintToken(cmd *cobra.Command, args []string){
	client := horizonclient.DefaultTestNetClient

	IssuerKeypair := keypair.MustParseFull(issuerSecret)
	DistributorKeypair := keypair.MustParseFull(distributorSecret)

	MintToken := <-service.MintToken(IssuerKeypair, DistributorKeypair.Address(), amount, tokenSymbol, client)
	log.Println("Mint Token Transaction Hash : ", MintToken)

}
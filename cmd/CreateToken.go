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
	createTokenCmd := &cobra.Command{
		Use:   "CreateToken",
		Short: "Create a new Stellar token",
		Run:   createToken,
	}

	// Add the flags for 'CreateToken' command
	createTokenCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the new token")
	createTokenCmd.MarkFlagRequired("name")
	createTokenCmd.Flags().StringVarP(&issuerSecret, "issuer", "i", "", "The secret key of the issuer account")
	createTokenCmd.MarkFlagRequired("issuer")

	// Add the 'CreateToken' command to the root command
	rootCmd.AddCommand(createTokenCmd)
}

func createToken(cmd *cobra.Command, args []string) {
	log.Println("Issuing new Token and Distributor ...")

	client := horizonclient.DefaultTestNetClient
	IssuerKeypair := keypair.MustParseFull(issuerSecret)

	distributorKP, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	TransferXLMResponse := <-service.InitializeDefaultDistributorAccount(IssuerKeypair, distributorKP.Address(), client)

	log.Println("Transaction Hashes : ")
	log.Println("InitializeDefaultDistributorAccount Transaction Hash : ", TransferXLMResponse)

	ChangeTrustResponse := <-service.ChangeTrust(distributorKP, IssuerKeypair.Address(), name, "100000", client)

	log.Println("Issuing New Token Transaction Hash : ", ChangeTrustResponse)

	log.Println("Distributor Account Detail : ")
	log.Println("Distributor Address / Public Key : ", distributorKP.Address())
	log.Println("Distributor Secret / Seed Key : ", distributorKP.Seed())
	
	log.Println("User Input Arguments : ")
	log.Println("Token Name : ", name)
	log.Println("Issuer Secret / Seed Key : ", issuerSecret)


}


package main

import (
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestContract(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	println("private-key ", key.D.Text(16))

	conn := backends.NewSimulatedBackend(core.GenesisAlloc{
		auth.From: {
			PrivateKey: key.D.Bytes(),
			Balance:    big.NewInt(10000000000),
		},
	})

	address, tx, token, err := DeployFixedSupplyToken(auth, conn)

	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}

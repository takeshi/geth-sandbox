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
	"github.com/ethereum/go-ethereum/ethclient"
)

func testMain(t *testing.T, conn bind.ContractBackend, auth *bind.TransactOpts) {
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

func startNode() {

}

func TestByNode(t *testing.T) {

	privateKey, err := crypto.HexToECDSA("149c6afe729dc649a4bcce4d79a1c594b3f1fc8f503175132daf637830e05cec")
	// address := crypto.PubkeyToAddress(privateKey.PublicKey)

	conn, err := ethclient.Dial("../data/geth.ipc")
	if err != nil {
		t.Errorf("Failed to connect to the Ethereum client: %v", err)
	}

	// miner, err := bind.NewTransactor(strings.NewReader(key), pass)
	user := bind.NewKeyedTransactor(privateKey)
	// miner, err := bind.NewTransactor(strings.NewReader(key), pass)
	// user := bind.NewKeyedTransactor(privateKey)

	testMain(t, conn, user)

}

func TestContractBySimulator(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	println("private-key ", key.D.Text(16))

	conn := backends.NewSimulatedBackend(core.GenesisAlloc{
		auth.From: {
			PrivateKey: key.D.Bytes(),
			Balance:    big.NewInt(10000000000),
		},
	})

	testMain(t, conn, auth)
}

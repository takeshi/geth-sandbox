package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
)

const key = `
{"address":"25477d57c97737abb7789efc42edc412be0596e9","crypto":{"cipher":"aes-128-ctr","ciphertext":"b11bfcdb75bc9ba12ff04a8842a02dec0e90318e327bf0a9ad961a03af09c41b","cipherparams":{"iv":"bd1c6bf022b248e511e05f09fe33478a"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"40b8ca2296e5adf5b6eecd6601f503d22497594724a3964f387f2f5f8cc0e4ea"},"mac":"d8106c5a58480f615071de0865c8dd163d615a0754f4202e3d37d7b198032de8"},"id":"d6f5eb77-bda5-4f3e-a4e1-6226f295a9cb","version":3}
`
const pass = "hoge"

func waitTransaction(conn ethereum.TransactionReader, tra *types.Transaction, interval time.Duration, retry int) (*types.Receipt, error) {
	println("wait transaction  " + tra.Hash().Hex())
	ctx := context.Background()
	for i := 0; i < retry; i++ {
		receipt, err := conn.TransactionReceipt(ctx, tra.Hash())
		if err != nil {
			print(".")
			// print(err.Error())
			time.Sleep(interval)
			continue
		}
		return receipt, nil
	}
	return nil, errors.New("Retry Error Transaction Hash: " + tra.Hash().Hex())
}

func sendEth(conn bind.ContractTransactor, to common.Address, traOps *bind.TransactOpts) (*types.Transaction, error) {
	bound := bind.NewBoundContract(to, abi.ABI{}, nil, conn, nil)
	tx1, err := bound.Transfer(traOps)
	if err != nil {
		log.Fatalf("SendTransaction transaction error: %v", err)
		return nil, err
	}
	return tx1, nil
}

// func open() *ethclient.Client{

// }

func deploy() {

	// privateKey, err := crypto.GenerateKey()
	privateKey, err := crypto.HexToECDSA("149c6afe729dc649a4bcce4d79a1c594b3f1fc8f503175132daf637830e05cec")
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	conn, err := ethclient.Dial("../data/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	miner, err := bind.NewTransactor(strings.NewReader(key), pass)
	user := bind.NewKeyedTransactor(privateKey)

	println("user " + user.From.Hex())
	println("miner " + miner.From.Hex())

	// Deploy するための Gas を送りつける
	tx1, err := sendEth(conn, user.From, &bind.TransactOpts{
		From:     miner.From,
		Signer:   miner.Signer,
		Value:    big.NewInt(5 * (2 << 20)),
		GasLimit: 500000,
	})
	if err != nil {
		log.Fatalf("SendTransaction transaction error: %v", err)
	}
	println("-------------------------------")
	println("GasPrise ", tx1.GasPrice().Text(10))
	println("Gas ", tx1.Gas())

	receipt, err := waitTransaction(conn, tx1, 1500*time.Millisecond, 30)
	if err != nil {
		log.Fatalf("transaction timeout error: %v", err)
	}
	println("Status ", receipt.Status)
	println("GasUsed ", receipt.GasUsed)

	// GASの量を確認する
	ctx := context.TODO()
	blanceAuth, _ := conn.BalanceAt(ctx, user.From, nil)
	blanceMiner, _ := conn.BalanceAt(ctx, miner.From, nil)

	println("blance auth", blanceAuth.Text(10))
	println("balance miner", blanceMiner.Text(10))

	println("-------------------------------")

	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := DeployFixedSupplyToken(miner, conn) //, new(big.Int), "Contracts in Go!!!", 0, "Go!")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	receipt2, err := waitTransaction(conn, tx, 1500*time.Millisecond, 20)
	if err != nil {
		log.Fatalf("transaction timeout error: %v", err)
	}
	println("Status ", receipt2.Status)
	println("GasUsed ", receipt2.GasUsed)

	owner, err := token.Owner(&bind.CallOpts{
		From: miner.From,
	})
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	println("owner " + owner.Hex())

}

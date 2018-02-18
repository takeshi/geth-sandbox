package main

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/eth"
	cli "gopkg.in/urfave/cli.v1"
)

func TestMain(t *testing.T) {

	// create mock cli data
	set := createFlagSet()
	ctx := cli.NewContext(app, set, nil)

	// start geth
	node := makeFullNode(ctx)
	startNode(ctx, node)
	defer node.Stop()

	// start minining
	var eth *eth.Ethereum
	node.Service(&eth)

	if err := eth.StartMining(true); err != nil {
		t.Errorf("Failed to start mining: %v", err)
	}

	// check mining status
	time.Sleep(3 * time.Second)

	t.Log("isMining ", eth.IsMining())
	if !eth.IsMining() {
		t.Errorf("mining did't start")
	}

}

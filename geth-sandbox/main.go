// Copyright 2014 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// geth is the official command-line client for Ethereum.
package main

import (
	"flag"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"gopkg.in/urfave/cli.v1"
)

const (
	clientIdentifier = "geth" // Client identifier to advertise over the network
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	// Ethereum address of the Geth release oracle.
	relOracle = common.HexToAddress("0xfa7b9770ca4cb04296cac84f37736d4041251cdf")
	// The app that holds all commands and flags.
	app = utils.NewApp(gitCommit, "the go-ethereum command line interface")
	// flags that configure the node

)

func createFlagSet() *flag.FlagSet {
	set := flag.NewFlagSet("geth-sandbox", flag.ContinueOnError)
	utils.GCModeFlag.Apply(set)
	return set
}

func main() {
	set := createFlagSet()
	ctx := cli.NewContext(app, set, nil)

	node := makeFullNode(ctx)
	startNode(ctx, node)
	defer node.Stop()

	startMining(ctx, node)
	createConsole(ctx, node)

	node.Wait()
}

func startMining(ctx *cli.Context, node *node.Node) {
	var eth *eth.Ethereum
	node.Service(&eth)

	eth.TxPool().SetGasPrice(utils.GlobalBig(ctx, utils.GasPriceFlag.Name))
	if err := eth.StartMining(true); err != nil {
		utils.Fatalf("Failed to start mining: %v", err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		println("minning ", eth.IsMining())
	}()

}

func createConsole(ctx *cli.Context, node *node.Node) {
	client, err := node.Attach()
	config := console.Config{
		DataDir: "/my/tmp/data",
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}
	console, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}
	defer console.Stop(false)
	console.Welcome()
	console.Interactive()
}

func startNode(ctx *cli.Context, stack *node.Node) {
	utils.StartNode(stack)
}

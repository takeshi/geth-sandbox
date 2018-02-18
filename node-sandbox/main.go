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
	"gopkg.in/urfave/cli.v1"

	// "github.com/ethereum/go-ethereum/cmd/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/node"

	"github.com/takeshi/geth-sandbox/utils"

	sandbox "github.com/takeshi/geth-sandbox/sandbox"
)

const (
	clientIdentifier = "geth" // Client identifier to advertise over the network
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	// Ethereum address of the Geth release oracle.
	// relOracle = common.HexToAddress("0xfa7b9770ca4cb04296cac84f37736d4041251cdf")
	// The app that holds all commands and flags.
	app = utils.NewApp(gitCommit, "the go-ethereum command line interface")
	// flags that configure the node

)

func main() {
	privateKey, _ := crypto.HexToECDSA("149c6afe729dc649a4bcce4d79a1c594b3f1fc8f503175132daf637830e05cec")
	user := bind.NewKeyedTransactor(privateKey)

	set := sandbox.CreateFlagSet()

	ctx := cli.NewContext(app, set, nil)

	node := sandbox.MakeFullNode(ctx)
	startNode(ctx, node)
	defer node.Stop()

	sandbox.StartMining(ctx, node, &user.From)
	sandbox.CreateConsole(ctx, node)

	node.Wait()
}

func startNode(ctx *cli.Context, stack *node.Node) {
	utils.StartNode(stack)
}

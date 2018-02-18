package sandbox

import (
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/takeshi/geth-sandbox/utils"
	cli "gopkg.in/urfave/cli.v1"
)

func StartMining(ctx *cli.Context, node *node.Node, ethebase *common.Address) {
	var eth *eth.Ethereum
	node.Service(&eth)

	println("test coinbase ", strings.ToLower(ethebase.Hex()))
	eth.SetEtherbase(*ethebase)
	eth.Miner().SetEtherbase(*ethebase)
	// time.Sleep(1 * time.Second)

	// eth.TxPool().SetGasPrice(utils.GlobalBig(ctx, utils.GasPriceFlag.Name))

	// eth.
	type threaded interface {
		SetThreads(threads int)
	}
	if th, ok := eth.Engine().(threaded); ok {
		th.SetThreads(1)
	}

	if err := eth.StartMining(true); err != nil {
		utils.Fatalf("Failed to start mining: %v", err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		println("minning ", eth.IsMining())
	}()

}

func CreateConsole(ctx *cli.Context, node *node.Node) {
	client, err := node.Attach()
	config := console.Config{
		DataDir: "./data",
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

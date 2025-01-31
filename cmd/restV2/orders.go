package restV2

import (
	"github.com/wselfjes/bitget"
	"github.com/wselfjes/bitget/cmd/config"
)

func Orders() {
	println()
	rest := bitget.NewRest(config.FullApi())

	getPendingOrders(rest)

	println("\n")

	rest1 := bitget.NewRest(config.PereApi())

	getPendingOrders(rest1)
}

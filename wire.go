//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func initializeGarden(numFlowers int) *Garden {
	wire.Build(GrowGarden, GrowCaretaker, GrowFlowersRandom)
	return &Garden{}
}

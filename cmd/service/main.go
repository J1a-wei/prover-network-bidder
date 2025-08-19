package main

import (
	"github.com/brevis-network/prover-network-bidder/service"
	"github.com/celer-network/goutils/log"
)

func main() {
	err := service.Start()
	if err != nil {
		log.Errorln(err)
	}
}

package scheduler

import (
	"github.com/brevis-network/prover-network-bidder/client"
	"github.com/brevis-network/prover-network-bidder/dal"
	"github.com/brevis-network/prover-network-bidder/onchain"
)

type Scheduler struct {
	*dal.DAL
	*onchain.ChainClient
	*client.ProverNetworkClient
}

func NewScheduler(db *dal.DAL, c *onchain.ChainClient, p *client.ProverNetworkClient) *Scheduler {
	return &Scheduler{db, c, p}
}

func (s *Scheduler) Start() {
	go s.scheduleBid()
	go s.scheduleProove()
}

func (s *Scheduler) scheduleBid() {
	// TODO
}

func (s *Scheduler) scheduleProove() {
	// TODO
}

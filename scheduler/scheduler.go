package scheduler

import (
	"context"
	"time"

	"github.com/brevis-network/prover-network-bidder/client"
	"github.com/brevis-network/prover-network-bidder/dal"
	"github.com/brevis-network/prover-network-bidder/onchain"
	"github.com/celer-network/goutils/log"
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
	go s.scheduleAppRegister()
	go s.scheduleBid()
	go s.scheduleProve()
}

func (s *Scheduler) scheduleAppRegister() {
	for {
		time.Sleep(5 * time.Second)
		apps, err := s.FindNotRegisteredApps(context.Background())
		if err != nil {
			log.Errorf("FindNotRegisteredApps err: %s", err)
			continue
		}
		var elf []byte
		// TODO: retrieve elf from img url
		for _, app := range apps {
			err = s.RegisterApp(app.AppID, "", elf)
			if err != nil {
				log.Errorf("RegisterApp %s err: %s", app.AppID, err)
				continue
			}

			err = s.UpdateAppAsRegistered(context.Background(), app.AppID)
			if err != nil {
				log.Errorf("UpdateAppAsRegistered %s err: %s", app.AppID, err)
				continue
			}
		}
	}
}

func (s *Scheduler) scheduleBid() {
	for {
		time.Sleep(5 * time.Second)
		// TODO
	}
}

func (s *Scheduler) scheduleProve() {
	for {

		time.Sleep(5 * time.Second)
		// TODO
	}
}

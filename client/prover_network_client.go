package client

import (
	"context"
	"fmt"

	"github.com/brevis-network/prover-network-gateway/client/serviceapi"
	"github.com/celer-network/goutils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProverNetworkClient struct {
	conn *grpc.ClientConn
}

func NewProverNetworkClient(nodeUrl string) (*ProverNetworkClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(nodeUrl, opts...)
	if err != nil {
		return nil, err
	}
	mc := &ProverNetworkClient{
		conn: conn,
	}
	return mc, nil
}

func (c *ProverNetworkClient) RegisterApp(appId, info string, elf []byte) error {
	resp, err := serviceapi.NewProverNetworkClient(c.conn).RegisterApp(context.Background(), &serviceapi.RegisterAppRequest{
		AppId: appId,
		Elf:   elf,
		Info:  &info,
	})
	if err != nil {
		return err
	}
	if resp.GetErr() != nil {
		log.Errorf("RegisterApp err: %s", resp.GetErr().GetMsg())
		return fmt.Errorf("RegisterApp err: %s", resp.GetErr().GetMsg())
	}
	return nil
}

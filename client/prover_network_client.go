package client

import (
	"context"
	"fmt"

	"github.com/brevis-network/prover-network-bidder/client/serviceapi"
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

func (c *ProverNetworkClient) EstimateCost(appId string, inputs [][]byte) (cost uint32, pvDigest []byte, err error) {
	resp, err := serviceapi.NewProverNetworkClient(c.conn).EstimateCost(context.Background(), &serviceapi.EstimateCostRequest{
		AppId:  appId,
		Inputs: inputs,
	})
	if err != nil {
		return 0, []byte{}, err
	}
	if resp.GetErr() != nil {
		log.Errorf("EstimateCost err: %s", resp.GetErr().GetMsg())
		return 0, []byte{}, fmt.Errorf("EstimateCost err: %s", resp.GetErr().GetMsg())
	}
	return resp.Cost, resp.PvDigest, nil
}

func (c *ProverNetworkClient) ProveTask(appId, taskId string, inputs [][]byte) error {
	resp, err := serviceapi.NewProverNetworkClient(c.conn).ProveTask(context.Background(), &serviceapi.ProveTaskRequest{
		AppId:  appId,
		TaskId: taskId,
		Inputs: inputs,
	})
	if err != nil {
		return err
	}
	if resp.GetErr() != nil {
		log.Errorf("ProveTask err: %s", resp.GetErr().GetMsg())
		return fmt.Errorf("ProveTask err: %s", resp.GetErr().GetMsg())
	}
	return nil
}

func (c *ProverNetworkClient) GetProvingResult(appId, taskId string, inputs [][]byte) (proof []byte, err error) {
	resp, err := serviceapi.NewProverNetworkClient(c.conn).GetProvingResult(context.Background(), &serviceapi.GetProvingResultRequest{
		AppId:  appId,
		TaskId: taskId,
	})
	if err != nil {
		return []byte{}, err
	}
	if resp.GetErr() != nil {
		log.Errorf("GetProvingResult err: %s", resp.GetErr().GetMsg())
		return []byte{}, fmt.Errorf("GetProvingResult err: %s", resp.GetErr().GetMsg())
	}
	return resp.Proof, nil
}

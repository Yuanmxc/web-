package rpc

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/ticket"
	"TTMS/kitex_gen/ticket/ticketservice"
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var ticketClient ticketservice.Client

func InitTicketRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := ticketservice.NewClient(
		consts.TicketServiceName,
		//client.WithShortConnection(),
		client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 3, MaxIdlePerAddress: 100, MaxIdleGlobal: 10000, MaxIdleTimeout: time.Minute}),
		//client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(consts.RPCTimeout),          // rpc timeout
		client.WithConnectTimeout(consts.ConnectTimeout),  // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	ticketClient = c
}

func UpdateTicket(ctx context.Context, req *ticket.UpdateTicketRequest) (*ticket.UpdateTicketResponse, error) {
	resp, err := ticketClient.UpdateTicket(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetAllTicket(ctx context.Context, req *ticket.GetAllTicketRequest) (*ticket.GetAllTicketResponse, error) {
	resp, err := ticketClient.GetAllTicket(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func BuyTicket(ctx context.Context, req *ticket.BuyTicketRequest) (*ticket.BuyTicketResponse, error) {
	resp, err := ticketClient.BuyTicket(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func ReturnTicket(ctx context.Context, req *ticket.ReturnTicketRequest) (*ticket.ReturnTicketResponse, error) {
	resp, err := ticketClient.ReturnTicket(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

package rpc

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/order"
	"TTMS/kitex_gen/order/orderservice"
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var orderClient orderservice.Client

func InitOrderRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := orderservice.NewClient(
		consts.OrderServiceName,
		//client.WithShortConnection(),
		//client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 1, MaxIdlePerAddress: 100, MaxIdleGlobal: 10000, MaxIdleTimeout: time.Minute}),
		//client.WithMiddleware(mw.CommonMiddleware),
		//client.WithInstanceMW(mw.ClientMiddleware),
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
	orderClient = c
}

func GetAllOrder(ctx context.Context, req *order.GetAllOrderRequest) (*order.GetAllOrderResponse, error) {
	resp, err := orderClient.GetAllOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetOrderAnalysis(ctx context.Context, req *order.GetOrderAnalysisRequest) (*order.GetOrderAnalysisResponse, error) {
	resp, err := orderClient.GetOrderAnalysis(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func CommitOrder(ctx context.Context, req *order.CommitOrderRequest) (*order.CommitOrderResponse, error) {
	resp, err := orderClient.CommitOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

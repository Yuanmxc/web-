package rpc

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/play"
	"TTMS/kitex_gen/play/playservice"
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var playClient playservice.Client

func InitPlayRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := playservice.NewClient(
		consts.PlayServiceName,
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
	playClient = c
}

func AddPlay(ctx context.Context, req *play.AddPlayRequest) (*play.AddPlayResponse, error) {
	resp, err := playClient.AddPlay(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func UpdatePlay(ctx context.Context, req *play.UpdatePlayRequest) (*play.UpdatePlayResponse, error) {
	resp, err := playClient.UpdatePlay(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func DeletePlay(ctx context.Context, req *play.DeletePlayRequest) (*play.DeletePlayResponse, error) {
	resp, err := playClient.DeletePlay(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetAllPlay(ctx context.Context, req *play.GetAllPlayRequest) (*play.GetAllPlayResponse, error) {
	resp, err := playClient.GetAllPlay(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func AddSchedule(ctx context.Context, req *play.AddScheduleRequest) (*play.AddScheduleResponse, error) {
	resp, err := playClient.AddSchedule(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func UpdateSchedule(ctx context.Context, req *play.UpdateScheduleRequest) (*play.UpdateScheduleResponse, error) {
	resp, err := playClient.UpdateSchedule(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func DeleteSchedule(ctx context.Context, req *play.DeleteScheduleRequest) (*play.DeleteScheduleResponse, error) {
	resp, err := playClient.DeleteSchedule(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetAllSchedule(ctx context.Context, req *play.GetAllScheduleRequest) (*play.GetAllScheduleResponse, error) {
	resp, err := playClient.GetAllSchedule(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

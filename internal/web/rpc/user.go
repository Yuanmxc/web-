package rpc

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/user"
	"TTMS/kitex_gen/user/userservice"
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func InitUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		consts.UserServiceName,
		//client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 1, MaxIdlePerAddress: 100, MaxIdleGlobal: 10000, MaxIdleTimeout: time.Minute}),
		//client.WithMuxConnection(1),                       // 开启多路复用
		client.WithRPCTimeout(consts.RPCTimeout),          // rpc timeout
		client.WithConnectTimeout(consts.ConnectTimeout),  // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer 默认OpenTracing链路追踪
		client.WithResolver(r), // resolver
	)

	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func UserLogin(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetAllUser(ctx context.Context, req *user.GetAllUserRequest) (*user.GetAllUserResponse, error) {
	resp, err := userClient.GetAllUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func ChangeUserPassword(ctx context.Context, req *user.ChangeUserPasswordRequest) (*user.ChangeUserPasswordResponse, error) {
	resp, err := userClient.ChangeUserPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	resp, err := userClient.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	resp, err := userClient.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func BindEmail(ctx context.Context, req *user.BindEmailRequest) (*user.BindEmailResponse, error) {
	resp, err := userClient.BindEmail(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func ForgetPassword(ctx context.Context, req *user.ForgetPasswordRequest) (*user.ForgetPasswordResponse, error) {
	resp, err := userClient.ForgetPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	user "TTMS/kitex_gen/user"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, Req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error)
	UserLogin(ctx context.Context, Req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error)
	GetAllUser(ctx context.Context, Req *user.GetAllUserRequest, callOptions ...callopt.Option) (r *user.GetAllUserResponse, err error)
	ChangeUserPassword(ctx context.Context, Req *user.ChangeUserPasswordRequest, callOptions ...callopt.Option) (r *user.ChangeUserPasswordResponse, err error)
	DeleteUser(ctx context.Context, Req *user.DeleteUserRequest, callOptions ...callopt.Option) (r *user.DeleteUserResponse, err error)
	GetUserInfo(ctx context.Context, Req *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error)
	BindEmail(ctx context.Context, Req *user.BindEmailRequest, callOptions ...callopt.Option) (r *user.BindEmailResponse, err error)
	ForgetPassword(ctx context.Context, Req *user.ForgetPasswordRequest, callOptions ...callopt.Option) (r *user.ForgetPasswordResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, Req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, Req)
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, Req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, Req)
}

func (p *kUserServiceClient) GetAllUser(ctx context.Context, Req *user.GetAllUserRequest, callOptions ...callopt.Option) (r *user.GetAllUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllUser(ctx, Req)
}

func (p *kUserServiceClient) ChangeUserPassword(ctx context.Context, Req *user.ChangeUserPasswordRequest, callOptions ...callopt.Option) (r *user.ChangeUserPasswordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeUserPassword(ctx, Req)
}

func (p *kUserServiceClient) DeleteUser(ctx context.Context, Req *user.DeleteUserRequest, callOptions ...callopt.Option) (r *user.DeleteUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, Req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoRequest, callOptions ...callopt.Option) (r *user.GetUserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kUserServiceClient) BindEmail(ctx context.Context, Req *user.BindEmailRequest, callOptions ...callopt.Option) (r *user.BindEmailResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BindEmail(ctx, Req)
}

func (p *kUserServiceClient) ForgetPassword(ctx context.Context, Req *user.ForgetPasswordRequest, callOptions ...callopt.Option) (r *user.ForgetPasswordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ForgetPassword(ctx, Req)
}

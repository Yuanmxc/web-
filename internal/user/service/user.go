package service

import (
	"TTMS/internal/user/dao"
	"TTMS/kitex_gen/user"
	"context"
	"log"
)

func CreateUserService(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	userInfo := &user.User{
		Type:     req.Type,
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	err, id := dao.CreateUser(ctx, userInfo)
	resp := &user.CreateUserResponse{BaseResp: &user.BaseResp{}}
	resp.UserId = id
	if err != nil && err.Error() != "[警告] 添加成功，您设置了重名用户" {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else if err == nil {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "[警告] 添加成功，您设置了重名用户"
	}
	log.Println("resp = ", resp)
	return resp, nil
}
func UserLoginService(ctx context.Context, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	userInfo := &user.User{
		Id:       req.UserId,
		Password: req.Password,
	}
	u, err := dao.UserLogin(ctx, userInfo)
	resp := &user.UserLoginResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		resp.UserInfo = u
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
		resp.UserInfo = u
	}
	return resp, nil
}
func GetAllUserService(ctx context.Context, req *user.GetAllUserRequest) (*user.GetAllUserResponse, error) {
	resp := &user.GetAllUserResponse{BaseResp: &user.BaseResp{}, Data: &user.GetAllUserResponseData{}}
	var err error
	resp.Data.List, resp.Data.Total, err = dao.GetAllUser(ctx, int(req.Current), int(req.PageSize))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func ChangeUserPasswordService(ctx context.Context, req *user.ChangeUserPasswordRequest) (*user.ChangeUserPasswordResponse, error) {
	err := dao.ChangeUserPassword(ctx, int(req.UserId), req.Password, req.NewPassword)
	resp := &user.ChangeUserPasswordResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func DeleteUserService(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	err := dao.DeleteUser(ctx, int(req.UserId))
	resp := &user.DeleteUserResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func GetUserInfoService(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	u, err := dao.GetUserInfo(ctx, int(req.UserId))
	resp := &user.GetUserInfoResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	resp.User = u
	return resp, nil
}

func BindEmailService(ctx context.Context, req *user.BindEmailRequest) (*user.BindEmailResponse, error) {
	err := dao.BindEmail(ctx, req.UserId, req.GetEmail())
	resp := &user.BindEmailResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func ForgetPasswordService(ctx context.Context, req *user.ForgetPasswordRequest) (*user.ForgetPasswordResponse, error) {
	err := dao.ForgetPassword(ctx, req.Email, req.NewPassword)
	resp := &user.ForgetPasswordResponse{BaseResp: &user.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

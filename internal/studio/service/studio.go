package service

import (
	"TTMS/internal/studio/dao"
	"TTMS/kitex_gen/studio"
	"context"
)

func AddStudioService(ctx context.Context, req *studio.AddStudioRequest) (resp *studio.AddStudioResponse, err error) {
	err = dao.AddStudio(ctx, req.Name, req.RowsCount, req.ColsCount)
	resp = &studio.AddStudioResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func GetAllStudioService(ctx context.Context, req *studio.GetAllStudioRequest) (resp *studio.GetAllStudioResponse, err error) {
	studios, total, err := dao.GetAllStudio(ctx, int(req.Current), int(req.PageSize))
	resp = &studio.GetAllStudioResponse{BaseResp: &studio.BaseResp{}, Data: &studio.GetAllStudioResponseData{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	resp.Data.List = studios
	resp.Data.Total = total
	return resp, nil
}

func GetStudioService(ctx context.Context, req *studio.GetStudioRequest) (resp *studio.GetStudioResponse, err error) {
	s, err := dao.GetStudio(ctx, int(req.Id))
	resp = &studio.GetStudioResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	resp.Result = s
	return resp, nil
}

func UpdateStudioService(ctx context.Context, req *studio.UpdateStudioRequest) (resp *studio.UpdateStudioResponse, err error) {
	err = dao.UpdateStudio(ctx, &studio.Studio{Id: req.Id, Name: req.Name, RowsCount: req.RowsCount, ColsCount: req.ColsCount})
	resp = &studio.UpdateStudioResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func DeleteStudioService(ctx context.Context, req *studio.DeleteStudioRequest) (resp *studio.DeleteStudioResponse, err error) {
	err = dao.DeleteStudio(ctx, req.Id)
	resp = &studio.DeleteStudioResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

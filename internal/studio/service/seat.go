package service

import (
	"TTMS/internal/studio/dao"
	"TTMS/kitex_gen/studio"
	"context"
)

func AddSeatService(ctx context.Context, req *studio.AddSeatRequest) (*studio.AddSeatResponse, error) {
	seatInfo := &studio.Seat{
		StudioId: req.StudioId,
		Row:      req.Row,
		Col:      req.Col,
		Status:   req.Status,
	}
	err := dao.AddSeat(ctx, seatInfo)
	resp := &studio.AddSeatResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func GetAllSeatService(ctx context.Context, req *studio.GetAllSeatRequest) (*studio.GetAllSeatResponse, error) {
	seats, total, err := dao.GetAllSeat(ctx, int(req.StudioId), int(req.Current), int(req.PageSize))
	resp := &studio.GetAllSeatResponse{BaseResp: &studio.BaseResp{}, Data: &studio.GetAllSeatResponseData{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	resp.Data.List = seats
	resp.Data.Total = total
	return resp, nil
}
func UpdateSeatService(ctx context.Context, req *studio.UpdateSeatRequest) (*studio.UpdateSeatResponse, error) {
	seatInfo := &studio.Seat{
		StudioId: req.StudioId,
		Row:      req.Row,
		Col:      req.Col,
		Status:   req.Status,
	}
	err := dao.UpdateSeat(ctx, seatInfo)
	resp := &studio.UpdateSeatResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func DeleteSeatService(ctx context.Context, req *studio.DeleteSeatRequest) (*studio.DeleteSeatResponse, error) {
	seatInfo := &studio.Seat{
		StudioId: req.StudioId,
		Row:      req.Row,
		Col:      req.Col,
	}
	err := dao.DeleteSeat(ctx, seatInfo)
	resp := &studio.DeleteSeatResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

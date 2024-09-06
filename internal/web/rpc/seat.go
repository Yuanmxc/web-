package rpc

import (
	"TTMS/kitex_gen/studio"
	"context"
)

func AddSeat(ctx context.Context, req *studio.AddSeatRequest) (*studio.AddSeatResponse, error) {
	resp, err := studioClient.AddSeat(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetAllSeat(ctx context.Context, req *studio.GetAllSeatRequest) (*studio.GetAllSeatResponse, error) {
	resp, err := studioClient.GetAllSeat(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func UpdateSeat(ctx context.Context, req *studio.UpdateSeatRequest) (*studio.UpdateSeatResponse, error) {
	resp, err := studioClient.UpdateSeat(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func DeleteSeat(ctx context.Context, req *studio.DeleteSeatRequest) (*studio.DeleteSeatResponse, error) {
	resp, err := studioClient.DeleteSeat(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

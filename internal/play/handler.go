package main

import (
	"TTMS/internal/play/service"
	play "TTMS/kitex_gen/play"
	"context"
)

// PlayServiceImpl implements the last service interface defined in the IDL.
type PlayServiceImpl struct{}

// AddPlay implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) AddPlay(ctx context.Context, req *play.AddPlayRequest) (resp *play.AddPlayResponse, err error) {
	// TODO: Your code here...
	return service.AddPlayService(ctx, req)
}

// UpdatePlay implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) UpdatePlay(ctx context.Context, req *play.UpdatePlayRequest) (resp *play.UpdatePlayResponse, err error) {
	// TODO: Your code here...
	return service.UpdatePlayService(ctx, req)
}

// DeletePlay implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) DeletePlay(ctx context.Context, req *play.DeletePlayRequest) (resp *play.DeletePlayResponse, err error) {
	// TODO: Your code here...
	return service.DeletePlayService(ctx, req)
}

// GetAllPlay implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) GetAllPlay(ctx context.Context, req *play.GetAllPlayRequest) (resp *play.GetAllPlayResponse, err error) {
	// TODO: Your code here...
	return service.GetAllPlayService(ctx, req)
}

// AddSchedule implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) AddSchedule(ctx context.Context, req *play.AddScheduleRequest) (resp *play.AddScheduleResponse, err error) {
	// TODO: Your code here...
	return service.AddScheduleService(ctx, req)
}

// UpdateSchedule implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) UpdateSchedule(ctx context.Context, req *play.UpdateScheduleRequest) (resp *play.UpdateScheduleResponse, err error) {
	// TODO: Your code here...
	return service.UpdateScheduleService(ctx, req)
}

// DeleteSchedule implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) DeleteSchedule(ctx context.Context, req *play.DeleteScheduleRequest) (resp *play.DeleteScheduleResponse, err error) {
	// TODO: Your code here...
	return service.DeleteScheduleService(ctx, req)
}

// GetAllSchedule implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) GetAllSchedule(ctx context.Context, req *play.GetAllScheduleRequest) (resp *play.GetAllScheduleResponse, err error) {
	// TODO: Your code here...
	return service.GetAllScheduleService(ctx, req)
}

// PlayToSchedule implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) PlayToSchedule(ctx context.Context, req *play.PlayToScheduleRequest) (resp *play.PlayToScheduleResponse, err error) {
	// TODO: Your code here...
	return service.PlayToScheduleService(ctx, req)
}

// GetSchedule implements the PlayServiceImpl interface.
func (s *PlayServiceImpl) GetSchedule(ctx context.Context, req *play.GetScheduleRequest) (resp *play.GetScheduleResponse, err error) {
	// TODO: Your code here...
	return service.GetScheduleService(ctx, req)
}

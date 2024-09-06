package main

import (
	"TTMS/internal/studio/service"
	studio "TTMS/kitex_gen/studio"
	"context"
)

// StudioServiceImpl implements the last service interface defined in the IDL.
type StudioServiceImpl struct{}

// AddStudio implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) AddStudio(ctx context.Context, req *studio.AddStudioRequest) (resp *studio.AddStudioResponse, err error) {
	// TODO: Your code here...
	return service.AddStudioService(ctx, req)
}

// GetAllStudio implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) GetAllStudio(ctx context.Context, req *studio.GetAllStudioRequest) (resp *studio.GetAllStudioResponse, err error) {
	// TODO: Your code here...
	return service.GetAllStudioService(ctx, req)
}

// GetStudio implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) GetStudio(ctx context.Context, req *studio.GetStudioRequest) (resp *studio.GetStudioResponse, err error) {
	// TODO: Your code here...
	return service.GetStudioService(ctx, req)
}

// UpdateStudio implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) UpdateStudio(ctx context.Context, req *studio.UpdateStudioRequest) (resp *studio.UpdateStudioResponse, err error) {
	// TODO: Your code here...
	return service.UpdateStudioService(ctx, req)
}

// DeleteStudio implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) DeleteStudio(ctx context.Context, req *studio.DeleteStudioRequest) (resp *studio.DeleteStudioResponse, err error) {
	// TODO: Your code here...
	return service.DeleteStudioService(ctx, req)
}

// AddSeat implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) AddSeat(ctx context.Context, req *studio.AddSeatRequest) (resp *studio.AddSeatResponse, err error) {
	// TODO: Your code here...
	return service.AddSeatService(ctx, req)
}

// GetAllSeat implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) GetAllSeat(ctx context.Context, req *studio.GetAllSeatRequest) (resp *studio.GetAllSeatResponse, err error) {
	// TODO: Your code here...
	return service.GetAllSeatService(ctx, req)
}

// UpdateSeat implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) UpdateSeat(ctx context.Context, req *studio.UpdateSeatRequest) (resp *studio.UpdateSeatResponse, err error) {
	// TODO: Your code here...
	return service.UpdateSeatService(ctx, req)
}

// DeleteSeat implements the StudioServiceImpl interface.
func (s *StudioServiceImpl) DeleteSeat(ctx context.Context, req *studio.DeleteSeatRequest) (resp *studio.DeleteSeatResponse, err error) {
	// TODO: Your code here...
	return service.DeleteSeatService(ctx, req)
}

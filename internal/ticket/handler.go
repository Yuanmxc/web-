package main

import (
	"TTMS/internal/ticket/service"
	ticket "TTMS/kitex_gen/ticket"
	"context"
)

// TicketServiceImpl implements the last service interface defined in the IDL.
type TicketServiceImpl struct{}

// BatchAddTicket implements the TicketServiceImpl interface.
func (s *TicketServiceImpl) BatchAddTicket(ctx context.Context, req *ticket.BatchAddTicketRequest) (resp *ticket.BatchAddTicketResponse, err error) {
	// TODO: Your code here...
	return service.BatchAddTicketService(ctx, req)
}

// UpdateTicket implements the TicketServiceImpl interface.
func (s *TicketServiceImpl) UpdateTicket(ctx context.Context, req *ticket.UpdateTicketRequest) (resp *ticket.UpdateTicketResponse, err error) {
	// TODO: Your code here...
	return service.UpdateTicketService(ctx, req)
}

// GetAllTicket implements the TicketServiceImpl interface.
func (s *TicketServiceImpl) GetAllTicket(ctx context.Context, req *ticket.GetAllTicketRequest) (resp *ticket.GetAllTicketResponse, err error) {
	// TODO: Your code here...
	return service.GetAllTicketService(ctx, req)
}

// BuyTicket implements the TicketServiceImpl interface.
func (s *TicketServiceImpl) BuyTicket(ctx context.Context, req *ticket.BuyTicketRequest) (resp *ticket.BuyTicketResponse, err error) {
	// TODO: Your code here...
	return service.BuyTicketService(ctx, req)
}

// ReturnTicket implements the TicketServiceImpl interface.
func (s *TicketServiceImpl) ReturnTicket(ctx context.Context, req *ticket.ReturnTicketRequest) (resp *ticket.ReturnTicketResponse, err error) {
	// TODO: Your code here...
	return service.ReturnTicketService(ctx, req)
}

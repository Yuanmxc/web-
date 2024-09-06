package main

import (
	"TTMS/internal/order/service"
	order "TTMS/kitex_gen/order"
	"context"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// GetAllOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetAllOrder(ctx context.Context, req *order.GetAllOrderRequest) (resp *order.GetAllOrderResponse, err error) {
	// TODO: Your code here...
	return service.GetAllOrderService(ctx, req)
}

// GetOrderAnalysis implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderAnalysis(ctx context.Context, req *order.GetOrderAnalysisRequest) (resp *order.GetOrderAnalysisResponse, err error) {
	// TODO: Your code here...
	return service.GetOrderAnalysisService(ctx, req)
}

// CommitOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CommitOrder(ctx context.Context, req *order.CommitOrderRequest) (resp *order.CommitOrderResponse, err error) {
	// TODO: Your code here...
	return service.CommitOrderService(ctx, req)
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderRequest) (resp *order.UpdateOrderResponse, err error) {
	// TODO: Your code here...
	return service.UpdateOrderService(ctx, req)
}

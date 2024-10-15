package handler

import (
	"context"
	"order-management/services/common/genproto/orders"
	"order-management/services/orders/types"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {

	//service injection
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
	//register other order service
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId:    42,
		CustomerId: 2,
		ProductId:  1,
		Quantity:   10,
	}

	err := h.ordersService.CreateOrder(ctx, order)

	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "Success",
	}

	return res, nil

}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {

	o := h.ordersService.GetOrders(ctx)

	res := &orders.GetOrdersResponse{
		Orders: o,
	}

	return res, nil

}

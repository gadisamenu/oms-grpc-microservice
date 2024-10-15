package handler

import (
	"net/http"
	"order-management/services/common/genproto/orders"
	"order-management/services/common/util"
	"order-management/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var req orders.CreateOrderRequest

	err := util.ParseJSON(r, &req)

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)

		return
	}

	order := &orders.Order{
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
		OrderId:    42,
	}
	err = h.ordersService.CreateOrder(r.Context(), order)

	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	util.WriteJSON(w, http.StatusOK, res)
}

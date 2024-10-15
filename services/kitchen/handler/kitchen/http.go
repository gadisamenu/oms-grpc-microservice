package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"
	addrs "order-management/services/common/constants"
	"order-management/services/common/genproto/orders"
	"order-management/services/kitchen/types"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn

}

type HttpKitchenHandler struct {
	kitchenService types.KitchenService
}

func NewHttpKitchenHandler(kitchenService types.KitchenService) *HttpKitchenHandler {
	handler := &HttpKitchenHandler{
		kitchenService: kitchenService,
	}

	return handler
}

func (h *HttpKitchenHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("GET /", h.GetOrders)
}

func (h *HttpKitchenHandler) GetOrders(w http.ResponseWriter, r *http.Request) {

	conn := NewGRPCClient(addrs.ORDERS_GRPC_ADDR)
	defer conn.Close()

	c := orders.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
		CustomerId: 2,
		ProductId:  1,
		Quantity:   3,
	})

	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{CustomerId: 2})

	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	t := template.Must(template.New("orders").Parse(ordersTemplate))

	if err := t.Execute(w, res.Orders); err != nil {
		log.Fatalf("template error: %v", err)
	}

}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderId}}</td>
            <td>{{.CustomerId}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`

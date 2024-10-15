package main

import addrs "order-management/services/common/constants"

func main() {

	httpServer := NewHttpServer(addrs.ORDERs_HTTP_ADDR)
	go httpServer.Run()

	grpcServer := NewGRPCServer(addrs.ORDERS_GRPC_ADDR)
	grpcServer.Run()

}

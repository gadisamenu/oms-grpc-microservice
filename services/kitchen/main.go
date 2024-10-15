package main

import (
	"log"
	addrs "order-management/services/common/constants"
)

func main() {
	httpServer := NewHttpServer(addrs.KITCHEN_HTTP_ADDR)
	err := httpServer.Run()

	if err != nil {
		log.Fatalf("Kitchen Server failed: %v", err)
	}
}

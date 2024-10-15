package main

import (
	"log"
	"net/http"
	handler "order-management/services/kitchen/handler/kitchen"
	"order-management/services/kitchen/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	service := service.NewKitchenService()

	handler := handler.NewHttpKitchenHandler(service)
	handler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

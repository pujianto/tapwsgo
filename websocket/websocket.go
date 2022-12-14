package websocket

import (
	"log"
)

type webSocketService struct {
	listenAddr string
	endpoint   string
	name       string
}

func New(listenAddr string, endpoint string, name string) *webSocketService {
	ws := webSocketService{}
	ws.name = name
	ws.listenAddr = listenAddr
	ws.endpoint = endpoint
	return &ws
}

func htons(port int) int {
	return (port&0xff)<<8 | port>>8
}

func (w *webSocketService) Start() error {

	log.Printf("Starting %s service on %s", w.Name(), w.listenAddr)
	return nil
}

func (w *webSocketService) Stop() error {
	log.Printf("Stopping %s service on %s", w.Name(), w.listenAddr)
	return nil
}

func (w webSocketService) Name() string {
	return w.name
}

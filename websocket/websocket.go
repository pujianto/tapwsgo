package websocket

import (
	"log"

	"github.com/songgao/water"
)

type webSocketService struct {
	listenAddr string
	endpoint   string
	name       string
	iface      *water.Interface
}

func New(listenAddr string, endpoint string, name string, iface water.Interface) *webSocketService {
	ws := webSocketService{}
	ws.iface = &iface
	ws.name = name
	ws.listenAddr = listenAddr
	ws.endpoint = endpoint
	return &ws
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

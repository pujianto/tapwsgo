package main

import (
	"github.com/pujianto/tapwsgo/dhcp"
	"github.com/pujianto/tapwsgo/server"
	"github.com/pujianto/tapwsgo/websocket"
)

func main() {

	conf := server.LoadConfigFromEnv()
	iface := conf.GetInterface()
	var svcs = make([]server.Service, 0)
	ws1 := websocket.New("1.2.3.4:8080", "/ws", "websocket1", *iface)

	svcs = append(svcs, ws1)
	dhcpService := dhcp.New()
	svcs = append(svcs, dhcpService)

	s := server.New(conf, svcs...)
	s.Start()
	s.Stop()
}

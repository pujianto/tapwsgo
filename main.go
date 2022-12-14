package main

import (
	"github.com/pujianto/tapwsgo/dhcp"
	"github.com/pujianto/tapwsgo/server"
	"github.com/pujianto/tapwsgo/websocket"
)

func main() {
	//create a tap device
	// tap := &netlink.Tuntap{LinkAttrs: netlink.LinkAttrs{Name: "tap0"}, Mode: netlink.TUNTAP_MODE_TAP}
	// netlink.LinkAdd(tap)

	// //assign an IP address to the tap device
	// addr, err := netlink.ParseAddr("10.11.12.1/24")
	// netlink.AddrAdd(tap, addr)
	// fmt.Println(err)

	// //bring the tap device up
	// netlink.LinkSetUp(tap)

	// fmt.Println(tap)
	conf := server.LoadConfigFromEnv()
	var svcs = make([]server.Service, 0)
	ws1 := websocket.New("1.2.3.4:8080", "/ws", "websocket1")
	ws2 := websocket.New("1.2.3.4:8081", "/ws", "websocket2")

	svcs = append(svcs, ws1)
	svcs = append(svcs, ws2)
	dhcpService := dhcp.New()
	svcs = append(svcs, dhcpService)

	s := server.New(conf, svcs...)
	s.Start()
	s.Stop()
}

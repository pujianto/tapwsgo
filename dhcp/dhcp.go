package dhcp

type dhcpService struct {
	name string
}

func (d *dhcpService) Start() error {
	return nil
}

func (d *dhcpService) Stop() error {
	return nil
}

func (d *dhcpService) Name() string {
	return d.name
}

func New() *dhcpService {
	return &dhcpService{name: "dhcp"}
}

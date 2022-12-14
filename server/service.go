package server

type Service interface {
	Start() error
	Stop() error
	Name() string
}

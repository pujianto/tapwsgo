package server

import (
	"log"
)

type server struct {
	services []Service
	config   config
}

func New(config config, svcs ...Service) *server {
	s := server{config: config}
	s.services = svcs
	return &s
}

func (s *server) Start() error {
	for _, service := range s.services {
		if err := service.Start(); err != nil {
			log.Printf("Error starting service: %s", err)
			continue
		}
		log.Printf("Service started: %s", service.Name())
	}
	return nil
}

func (s *server) Stop() error {
	for _, service := range s.services {
		if err := service.Stop(); err != nil {
			log.Printf("Error stopping service: %s", err)
			continue
		}
		log.Printf("Service stopped: %s", service.Name())
	}
	return nil
}

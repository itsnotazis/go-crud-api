package api

import (
	"github.com/gofiber/fiber/v2"
)

type APIServer struct {
	Server *fiber.App
	Config *APIConfig
	Addr   string
}

type APIConfig struct {
	AppName       string
	ServerHeader  string
	Prefork       bool
	CaseSensitive bool
	StrictRouting bool
}

func NewAPIServer(f *fiber.App, config *APIConfig, addr string) (*APIServer, error) {
	conf := &APIConfig{
		AppName:       config.AppName,
		ServerHeader:  config.ServerHeader,
		Prefork:       config.Prefork,
		CaseSensitive: config.CaseSensitive,
		StrictRouting: config.StrictRouting,
	}
	return &APIServer{
		Server: f,
		Config: conf,
		Addr:   addr,
	}, nil
}

func (s *APIServer) Listen() error {
	return s.Server.Listen(s.Addr)
}

package server

import (
	"errors"
	"log"
	"os"
	"strings"
)

type config struct {
	listenAddr    string
	interfaceName string
	port          string
	interfaceIP   string
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (c config) validate() error {
	errs := make([]string, 0)
	if c.listenAddr == "" {
		errs = append(errs, "listen address is required")
	}
	if c.interfaceName == "" {
		errs = append(errs, "interface name is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
}

func LoadConfigFromEnv() config {
	defaults := map[string]string{
		"LOG_LEVEL":             "info",
		"HOST":                  "0.0.0.0",
		"PORT":                  "8080",
		"WITH_TLS":              "false",
		"TLS_CERT":              "/app/certs/fullchain.pem",
		"TLS_KEY":               "/app/certs/privkey.pem",
		"TLS_PASSPHRASE":        "",
		"WITH_DHCP":             "true",
		"INTERFACE_NAME":        "tapx",
		"INTERFACE_IP":          "10.11.12.254/24",
		"PUBLIC_INTERFACE_NAME": "eth0",
	}

	conf := config{}
	conf.listenAddr = getenv("HOST", defaults["HOST"])
	conf.port = getenv("PORT", defaults["PORT"])
	conf.interfaceName = getenv("INTERFACE_NAME", defaults["INTERFACE_NAME"])
	conf.interfaceIP = getenv("INTERFACE_IP", defaults["INTERFACE_IP"])

	if err := conf.validate(); err != nil {
		log.Fatalf("Invalid configurations: %s", err)
	}

	return conf
}

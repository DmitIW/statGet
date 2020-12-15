package tConnector

import (
	"github.com/tarantool/go-tarantool"
	"log"
)

var (
	defaultConnection *tarantool.Connection
)

func DefConnInit(addr string, cfg *tarantool.Opts) {
	if cfg == nil {
		cfg = getDefaultConfig()
	}
	defaultConnection = open(addr, *cfg)
}

func DefConnClose() {
	if err := defaultConnection.Close(); err != nil {
		log.Printf("Error on default connection closing: %v\n", err)
	}
}

func GetDefaultConnection() *tarantool.Connection {
	if defaultConnection == nil {
		log.Fatalf("Default connection is nil")
	}
	return defaultConnection
}

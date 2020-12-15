package tConnector

import (
	"github.com/tarantool/go-tarantool"
	"log"
)

func getDefaultConfig() *tarantool.Opts {
	return &tarantool.Opts{}
}

func open(addr string, opts tarantool.Opts) *tarantool.Connection {
	var (
		result *tarantool.Connection
		err    error
	)
	log.Printf("Open connection with tarantool on %v\n", addr)
	if result, err = tarantool.Connect(addr, opts); err != nil {
		log.Fatalf("Conection opening error: %v\n", err)
	}
	if resp, err := result.Ping(); err != nil {
		log.Fatalf("Tarantool connection is failed: %v\n", err)
	} else {
		log.Printf("Tarantool connection ping response: code = %v\n", resp.Code)
	}
	return result
}

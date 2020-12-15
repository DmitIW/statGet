package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	pathToEnvFile = ".env"
)

func LoadEnvFile() error {
	if err := godotenv.Load(pathToEnvFile); err != nil {
		return err
	}
	return nil
}

func getServerPort() string {
	if result := os.Getenv("SERVER_PORT"); len(result) != 0 {
		return result
	}
	return "4444"
}

func getTarantoolPort() string {
	if result := os.Getenv("TARANTOOL_PORT"); len(result) != 0 {
		return result
	}
	return "3301"
}

func getServerAddr() string {
	if result := os.Getenv("ADDR"); len(result) != 0 {
		return result
	}
	return ""
}

func getTarantoolAddr() string {
	if result := os.Getenv("TARANTOOL_ADDR"); len(result) != 0 {
		return result
	}
	return "127.0.0.1"
}

func concatAddr(addr string, port string) string {
	return addr + ":" + port
}

func GetSAddr() string {
	addr := getServerAddr()
	port := getServerPort()
	return concatAddr(addr, port)
}

func GetTAddr() string {
	addr := getTarantoolAddr()
	port := getTarantoolPort()
	return concatAddr(addr, port)
}

package main

import (
	"errors"
	"os"
	"strconv"
)

const DEFAULT_PORT = 8080

type Config struct {
	Port int
}

func NewConfig(file string) (*Config, error) {
	port, err := getPortFromEnv()

	if err != nil {
		return nil, err
	}

	port, err = checkPortValue(port)

	if err != nil {
		return nil, err
	}

	return &Config{Port: port}, nil
}

func getPortFromEnv() (int, error) {
	sPort := os.Getenv("PORT")

	if sPort == "" {
		return -1, errors.New("Port not found in env")
	}

	port, err := strconv.Atoi(sPort)

	if err != nil {
		return -1, err
	}

	return checkPortValue(port)

}

func checkPortValue(port int) (int, error) {
	if port < 0 || port > 65535 {
		return -1, errors.New("Port out of range")
	}

	return port, nil
}

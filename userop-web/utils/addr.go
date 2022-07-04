package utils

import (
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer lis.Close()
	return lis.Addr().(*net.TCPAddr).Port, nil
}

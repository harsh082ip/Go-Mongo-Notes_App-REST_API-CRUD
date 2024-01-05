package main

import (
	"fmt"
	"net"
	"testing"
)

func TestMongo(t *testing.T) {

	msg := checkPortAvailability(WEBPORT)
	if msg != "no error" {
		t.Errorf("Port %v is in use", WEBPORT)
	}
}

func checkPortAvailability(port string) string {

	conn, err := net.Dial("tcp", "localhost"+port)
	if err != nil {
		fmt.Printf("Port %v is available \n", port)
		return "no error"
	}
	defer conn.Close()

	return "Port is in use"
}

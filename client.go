package main

import (
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "1055"
	TYPE = "udp"
)

func main() {
	udpServer, err := net.ResolveUDPAddr(TYPE, HOST+":"+PORT)
	log.Println("Attempt Connect to:", HOST+":"+PORT, TYPE)

	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	_, err = conn.Write([]byte("{\"s\":null,\"e\":0,\"d\":{\"u\":\"\",\"s\":\"\"},\"r\":null}"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println(string(received))
}

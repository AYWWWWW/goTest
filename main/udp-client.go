package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var buf [512]byte
	service := "127.0.0.1:5000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkErr(err)
	defer conn.Close()
	rAddr := conn.RemoteAddr()
	n, err := conn.Write([]byte("Hello Server!"))
	checkErr(err)
	n, err = conn.Read(buf[0:])
	checkErr(err)
	fmt.Println("Reply from udp server ", rAddr.String(), string(buf[0:n]))
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var buf [512]byte
	service := "10.128.106.20:5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	defer conn.Close()
	rAddr := conn.RemoteAddr()
	n, err := conn.Write([]byte("Hello Server!"))
	checkErr(err)
	time.Sleep(2000)
	n, err = conn.Read(buf[0:])
	checkErr(err)
	fmt.Println("Reply from tcp server ", rAddr.String(), string(buf[0:n]))
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

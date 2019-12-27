package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func main() {
	fmt.Println("starting server on localhost:8081")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		panic(err)
	}

	originConn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer originConn.Close()

	fmt.Printf("proxying http request (%v) to origin server\n", req)
	err = req.Write(originConn)
	if err != nil {
		panic(err)
	}

	resp, err := http.ReadResponse(bufio.NewReader(originConn), req)
	if err != nil {
		panic(err)
	}

	resp.Write(conn)
}

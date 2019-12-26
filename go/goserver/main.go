package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
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
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		panic(err)
	}

	resp := &http.Response{
		Proto: "HTTP/1.1", ProtoMajor: 1, ProtoMinor: 1,
		Request: req, Header: make(http.Header, 0),
	}

	filepath := "." + req.URL.Path
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		resp.StatusCode = 400
		resp.Status = "400 Bad Request"

		resp.Body = ioutil.NopCloser(bytes.NewBufferString(err.Error()))
		resp.ContentLength = int64(len(err.Error()))
	} else {
		resp.StatusCode = 200
		resp.Status = "200 OK"

		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bs))
		resp.ContentLength = int64(len(bs))
	}

	err = resp.Write(conn)
	if err != nil {
		panic(err)
	}
}

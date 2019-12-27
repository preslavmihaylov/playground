package main

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func startTLS(conn net.Conn, connReader *bufio.Reader, server string) (net.Conn, *bufio.Reader) {
	sendCmd(conn, connReader, 220, "STARTTLS\n")
	config := &tls.Config{ServerName: server}
	conn = tls.Client(conn, config)
	fmt.Println("[INFO] TLS established")

	return conn, bufio.NewReader(conn)
}

func sendCmd(conn net.Conn, connReader *bufio.Reader, expected int, cmd string) string {
	if cmd != "" {
		fmt.Print("[CLIENT] ", cmd)
	}

	_, err := conn.Write([]byte(cmd))
	if err != nil {
		panic(err)
	}

	resp, err := connReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("[SERVER] ", resp)
	if getStatus(resp) != expected {
		panic("got unsuccessful status from server")
	}

	return resp
}

func getStatus(resp string) int {
	status, err := strconv.Atoi(strings.Split(resp, " ")[0])
	if err != nil {
		panic(err)
	}

	return status
}

func base64Encoded(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

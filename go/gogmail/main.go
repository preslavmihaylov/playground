package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
)

var server, port = "smtp.gmail.com", 587
var appPass string

func main() {
	readAppPass()
	sendEmail("from_email", "to_email", "Some subject", "Hi There")
}

func readAppPass() {
	bs, err := ioutil.ReadFile(".app_pass")
	if err != nil {
		panic(err)
	}

	appPass = string(bs)
}

func sendEmail(from, to, subject, body string) {
	conn, connReader := initialHandshake()
	defer conn.Close()

	auth(conn, connReader, from, appPass)
	mailFromTo(conn, connReader, from, to)
	data(conn, connReader, from, to, subject, body)
	quit(conn, connReader)
}

func initialHandshake() (net.Conn, *bufio.Reader) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		panic(err)
	}

	connReader := bufio.NewReader(conn)

	sendCmd(conn, connReader, 220, "")
	sendCmd(conn, connReader, 250, "HELO gmail.com\n")

	return startTLS(conn, connReader, server)
}

func auth(conn net.Conn, connReader *bufio.Reader, username, pass string) {
	sendCmd(conn, connReader, 334, "AUTH LOGIN\n")
	sendCmd(conn, connReader, 334, base64Encoded(username)+"\n")
	sendCmd(conn, connReader, 235, base64Encoded(pass)+"\n")
}

func mailFromTo(conn net.Conn, connReader *bufio.Reader, from, to string) {
	sendCmd(conn, connReader, 250, fmt.Sprintf("MAIL FROM: <%s>\n", from))
	sendCmd(conn, connReader, 250, fmt.Sprintf("RCPT TO: <%s>\n", to))
}

func data(conn net.Conn, connReader *bufio.Reader, from, to, subject, body string) {
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	sendCmd(conn, connReader, 354, "DATA\n")
	sendCmd(conn, connReader, 250, fmt.Sprintf("%s\r\n.\r\n", msg))
}

func quit(conn net.Conn, connReader *bufio.Reader) {
	sendCmd(conn, connReader, 221, "QUIT\n")
}

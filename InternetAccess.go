package main

import (
	"io"
	"os"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.google.co.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: www.google.co.jp \r\n\r\n"))
	io.Copy(os.Stdout, conn)
}
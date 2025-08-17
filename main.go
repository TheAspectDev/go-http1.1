package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := ls.Accept()
		go readConn(conn)

	}
}

func readConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	status, _, _ := reader.ReadLine()
	statusString := strings.TrimSpace(strings.Split(string(status), "/")[0])
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		if string(data) == "" {
			break
		}

		fmt.Printf("%s\n", string(data))

	}
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n"+
		"Content-Type: text/html\r\n"+
		"Server: Aspect\r\n\r\n"+
		"<h1>Request Type: %v</h1>", statusString)

}

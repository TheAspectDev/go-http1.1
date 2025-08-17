package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	rd, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	_, err = rd.Write([]byte("hi\n"))
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(rd)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("%s\n", data)
	}

}

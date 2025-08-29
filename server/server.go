package main

import (
	"fmt"
	"net"
	"strings"
)

var dict = map[string]string{
	"red": "красный",
	"green": "зеленый",
	"blue": "голубой",
	"yellow": "желтый",
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, 1024 * 4)
		n, err := conn.Read(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		source := strings.TrimSpace(string(input[:n]))
		source = strings.ToLower(source)
		target, ok := dict[source]
		if !ok {
			target = "undefined"
		}
		fmt.Println(source, "-", target)
		conn.Write([]byte(target))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening ...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err = conn.Write([]byte(source))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Перевод: ")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buff[:n]))
	}
}
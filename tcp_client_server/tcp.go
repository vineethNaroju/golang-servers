package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

/*
net.Dial and net.Listen return types implement io.Reader, io.Writer
*/

func TcpClient() {

	conn, err := net.Dial("tcp", "127.0.0.1:3421")

	if err != nil {
		log.Fatal(err)
	}

	arr := []string{"hello", "world", "noobs"}

	buf := make([]byte, 1024)

	rev := []string{}

	for _, str := range arr {
		conn.Write([]byte(str))

		cnt, err := conn.Read(buf)

		if err != nil {
			fmt.Println(str, err)
			continue
		}

		rev = append(rev, string(buf[:cnt]))
	}

	fmt.Println(rev)
}

func TcpServer() {

	server, err := net.Listen("tcp", ":3421")

	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	client, err := server.Accept()

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)

	for {
		cnt, err := client.Read(buf)

		if err != nil {
			log.Fatal(err)
		}

		lo, hi := 0, cnt-1

		for lo < hi {
			buf[lo], buf[hi] = buf[hi], buf[lo]
			lo++
			hi--
		}

		client.Write(buf[:cnt])
	}
}

func main() {
	go TcpServer()

	time.Sleep(2 * time.Second)

	go TcpClient()

	time.Sleep(5 * time.Second)
}

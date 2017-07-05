package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type Client struct {
	name string
	conn net.Conn
}

type ExStdin struct {
	prefix string
}

func (s *ExStdin) Read(p []byte) (n int, err error) {
	r := bytes.NewReader([]byte(s.prefix))
	io.Copy(os.Stdout, r)
	return os.Stdin.Read(p)
}

var clients = []*Client{}

func serve() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer ln.Close()
	log.Printf("Listening port 8000 ...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		newCl := &Client{
			name: strconv.Itoa(len(clients) + 1),
			conn: conn}
		clients = append(clients, newCl)

		go handleConn(newCl)
	}
}

func handleConn(client *Client) {
	conn := client.conn
	input := bufio.NewScanner(conn)
	fmt.Println("New user connected:", client.name)
	for input.Scan() {
		time.Sleep(1 * time.Second)
		for _, c := range clients {
			if client.name != c.name {
				c.conn.Write([]byte("\t" + client.name + ":" + input.Text() + "\n"))
			}
		}
	}
	conn.Close()
	fmt.Println("user exited:", client.name)
}

func conn() {
	c, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatalln("Failed to connect: ", err)
	}

	go io.Copy(c, &ExStdin{prefix: "me:"})
	io.Copy(os.Stdout, c)
}

func main() {
	switch os.Args[1] {
	case "serve":
		serve()
	case "conn":
		conn()
	default:
		log.Println("failed")
	}
}

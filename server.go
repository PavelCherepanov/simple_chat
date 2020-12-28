package main

import (
	"log"
	"net"
	"os"
	"time"
)

const (
	PROTOCOL = "tcp"
	PORT     = "127.0.0.1:808"
	BUFF     = 256
	QUAN     = 2
)

var connections = make(map[net.Conn]bool)

func main() {
	listen, err := net.Listen(PROTOCOL, PORT)
	if err != nil {
		os.Exit(1)
	}
	defer listen.Close()
	log.Println("[Server is listening]")
	for {
		if len(connections) < QUAN {
			conn, err := listen.Accept()
			if err != nil {
				break
			}
			tryConnection(conn)
			go connection(conn)
			//AddUser(conn)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func tryConnection(conn net.Conn) {
	conn.Write([]byte{1})
}

func connection(conn net.Conn) {
	defer conn.Close()
	connections[conn] = true
	var buffer []byte = make([]byte, BUFF)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			break
		}

		log.Print(string(buffer[:length]))
		for user := range connections {
			if user != conn {
				user.Write(buffer[:length])
			}
		}
	}

	delete(connections, conn)
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	PROTOCOL = "tcp"
	PORT     = "127.0.0.1:808"
	BUFF     = 256
)

var username string
logs := []string{}
func main() {
	conn, err := net.Dial(PROTOCOL, PORT)
	if err != nil {
		errorConnection()
		os.Exit(1)
	}
	defer conn.Close()

	tryConnection(conn)
	addUser(conn)
	go client(conn)
	readMessage(conn)
}

func tryConnection(conn net.Conn) {
	var check = make([]byte, 1)
	var timer = time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("Connection failure")
		os.Exit(1)
	}()
	conn.Read(check)
	timer.Stop()
	fmt.Println("Connection success")
}

func client(conn net.Conn) {
	for {
		fmt.Println("Success " + username)
		fmt.Println("Write hello to everyone")
		var message = inputString("")

		if len(message) != 0 {
			conn.Write([]byte(
				fmt.Sprintf("[%s]: %\n", username, message),
			))
		}
	}
}

func readMessage(conn net.Conn) {
	var buffer = make([]byte, BUFF)
	for {
		length, _ := conn.Read(buffer)
		if length != 0 {
			fmt.Print(string(buffer[:length]))
		}
	}
}

func inputString(text string) string {
	fmt.Print(text)
	message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(message, "\n", "", -1)
}

func errorConnection() {
	fmt.Println("-----ERROR-----")
}

func addUser(conn net.Conn) {
	username = inputString("Nickname: ")
	conn.Write([]byte(fmt.Sprintf("[Added user]: %s\n ", username)))
}

// writing all logs to the file

func writeFile(){
	file, error := os.Create("logs.txt")
	if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
	}
	defer file.Close() 
    file.WriteString(text)
     
    
}
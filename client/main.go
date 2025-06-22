package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listenAddr := ""
	if len(os.Args) < 2 {
		listenAddr = "127.0.0.1:8084"
		fmt.Printf("No Server address provided. Defaulting to %s\n", listenAddr)
	} else {
		listenAddr = os.Args[1]
	}

	networkType := "tcp"

	listen, err := net.Dial(networkType, listenAddr)
	defer listen.Close()

	if err != nil {
		fmt.Println("There was an err", err)
		os.Exit(1)
	}

	fmt.Printf("Connected to the server: %v\n", listenAddr)

	go writeToServer(listen)
	go readFromServer(listen)
	select{}
}

func writeToServer(conn net.Conn) {
	var givenLine string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		givenLine = scanner.Text()

		if givenLine == "exit" || givenLine == "quit" {
			fmt.Println("Exiting Client....")
			os.Exit(0)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}

		givenLineInByte := []byte(givenLine)

		_, err := conn.Write(givenLineInByte)

		if err != nil {
			fmt.Printf("Error in writing: %v", err)
		}
	}
}

func readFromServer(conn net.Conn){
	buffer := make([]byte, 1024)

	for{
		n, err := conn.Read(buffer)

		if err != nil {
			if err == io.EOF {
				fmt.Println("\nServer Disconnected...")
			} else {
				fmt.Printf("Error occurred in reading: %v", err)
			}
			return
		}

		receivedMessage := string(buffer[:n])
		fmt.Printf("From Server: %v", receivedMessage)
	}
}

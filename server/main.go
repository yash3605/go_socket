package main

import (
	"fmt"
	"io"
	"net"
	"github.com/joho/godotenv"
	"os"
)

func main(){
	error := godotenv.Load("./.env")
	if error != nil {
		fmt.Printf("Cannot load env file: %v", error)
	}


	listenAddress := os.Getenv("SERVER_LISTEN_ADDRESS")
	networkType := "tcp"

	if listenAddress == "" {
		listenAddress = "127.0.0.1:8084" // Default for listening on all interfaces
		fmt.Printf("SERVER_LISTEN_ADDRESS environment variable not set. Defaulting to %s\n", listenAddress)
	}

	listener , err := net.Listen(networkType, listenAddress)
	if err != nil {
		fmt.Println("Something went wrong with the listener:", err)
		os.Exit(1)
	}

	fmt.Printf("Server is listening on Address: %v\n", listenAddress)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("Client connected from: %v\n", remoteAddr)

	buffer := make([]byte, 1024)
	for {
		reply := "Received Client's reply\n"
		byteReply := []byte(reply)

		n, readErr := conn.Read(buffer)
		if readErr != nil {
			if readErr == io.EOF {
				fmt.Printf("Client %v disconnected.\n", remoteAddr)
			} else {
				fmt.Printf("Error reading data from %v: %v\n", remoteAddr, readErr)
			}
			return
		}
		receivedData := string(buffer[:n])
		fmt.Printf("Received from %v: '%s'\n", remoteAddr, receivedData)

		_, err := conn.Write(byteReply)

		if err != nil {
			fmt.Printf("Error occurred in replying : %v", err)
		}
	}
}

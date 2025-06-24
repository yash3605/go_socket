# Go TCP Chat Application

This project implements a basic concurrent TCP (Transmission Control Protocol) chat application using Go, demonstrating fundamental concepts of network programming, including sockets, concurrency (goroutines), and environment variable configuration.

The application consists of two separate programs: a **Server** that listens for incoming client connections and handles messages, and a **Client** that connects to the server, sends user input, and receives server replies.

---

## Project Structure
go_socket/
├── server/
│   ├── main.go        # Server application source code
│   └── .env           # Environment variables for server configuration
├── client/
│   └── main.go        # Client application source code
├── go.mod             # Go module file
└── go.sum             # Go module checksums

---

## Features

*   **Concurrent Server:** The server can handle multiple client connections simultaneously using Go goroutines. Each connected client gets its own dedicated goroutine for communication.
*   **Bidirectional Communication:** Both the client and server can send and receive messages continuously over an established connection.
*   **Configurable Server Address:** The server's listening IP address and port can be configured via an environment variable (`SERVER_LISTEN_ADDRESS`). A default is provided if the variable is not set.
*   **Client Connection via Argument:** The client connects to a server address specified as a command-line argument, allowing flexibility in targeting different servers. A default local address is provided for convenience.
*   **Interactive Chat:** Users can type messages in the client terminal, which are sent to the server, and the server's replies are displayed back in the client.

---

## How to Run

Follow these steps to set up and run the chat application.

### 1. Initialize Go Module (if not already done)

Navigate to the root of your project (`go_socket/`) and initialize the Go module:

```bash
cd go_socket
go mod init go_socket # Or your desired module name
```

### 2. Install Dependencies

The server uses the `github.com/joho/dotenv` package to read environment variables from a `.env` file. Install it:

```bash
go get github.com/joho/dotenv
```

### 3. Configure the Server

Create `.env` file inside `server/` directory(`go_socket/server/.env`) and add the server's listening address.

`SERVER_LISTEN_ADDRESS="<Address>"`

### 4. Run the Server

Open your first terminal and navigate into the `server/` directory.

```bash
cd go_socket/server
go run main.go
```

### 5. Run the Client(s)

Open another terminal and navigate to the `client/` directory.

```bash
cd go_socket/client/
go run main.go <SERVER_LISTEN_ADDRESS>
```

### 6. Interact

** Type messages in the client terminal and press `Enter`.
** Observe the server terminal, which will show messages received from clients.
** Observe the client terminal, which will display the server's replies.
** You can run multiple client instances concurrently to see the server handle them all.
** To exit a client, type `exit` or `quit` and press `Enter`

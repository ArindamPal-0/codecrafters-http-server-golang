package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on http://127.0.0.1:4221")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("> Request")
	requestBytes := make([]byte, 1024)
	requestLength, err := conn.Read(requestBytes)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("> Request Length: ", requestLength)

	requestStr := string(requestBytes[:requestLength])
	fmt.Println(requestStr)

	lines := strings.Split(requestStr, "\r\n")

	// http request status line
	statusLineSplit := strings.Split(lines[0], " ")

	// http request path
	httpPath := statusLineSplit[1]

	var responseBytes []byte
	if httpPath == "/" {
		responseBytes = []byte("HTTP/1.1 200 OK\r\n\r\n")
	} else {
		responseBytes = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	}

	_, err = conn.Write(responseBytes)
	if err != nil {
		fmt.Println("Error sending response: ", err.Error())
		os.Exit(1)
	}

}

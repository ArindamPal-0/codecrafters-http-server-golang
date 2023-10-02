package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Request struct {
	Method      string
	Path        string
	HttpVersion string
	Headers     map[string]string
	Body        []byte
}

func ParseHeaders(headersStr []string) map[string]string {
	headers := map[string]string{}

	for _, headerStr := range headersStr {
		headerSplit := strings.Split(headerStr, ":")

		key := strings.ToLower(strings.TrimSpace(headerSplit[0]))
		value := strings.TrimSpace(headerSplit[1])

		headers[key] = value
	}

	return headers
}

func ParseRequest(requestStr string) Request {

	lines := strings.Split(requestStr, "\r\n")

	// parse the request start line
	// http request status line
	startLineSplit := strings.Split(lines[0], " ")

	// http method
	httpMethod := startLineSplit[0]

	// http request path
	httpPath := startLineSplit[1]

	// http version
	httpVersion := startLineSplit[2]

	// http headers

	// get the line index after http headers
	headersEndIndex := 0
	for index, line := range lines {
		if line == "" {
			headersEndIndex = index
			break
		}
	}
	headerLines := lines[1:headersEndIndex]
	headers := ParseHeaders(headerLines)

	// http body
	body := []byte{}

	return Request{
		Method:      httpMethod,
		Path:        httpPath,
		HttpVersion: httpVersion,
		Headers:     headers,
		Body:        body,
	}
}

type StatusCode string

const (
	Status_NotFound StatusCode = "404 Not Found"

	Status_OK = "200 OK"
)

type Response struct {
	HttpVersion string
	StatusCode  StatusCode
	Headers     map[string]string
	Body        []byte
}

func (response *Response) getHeadersBytes() []byte {
	headerBytes := []byte{}
	for k, v := range response.Headers {
		headerStr := fmt.Sprintf("%s: %s\r\n", k, v)
		headerBytes = append(headerBytes, []byte(headerStr)...)
	}

	return headerBytes
}

func (response *Response) getResponseBytes() []byte {
	// append http response status line
	statusLine := fmt.Sprintf("%s %s\r\n", response.HttpVersion, response.StatusCode)
	responseBytes := []byte(statusLine)

	// append http response headers
	headersBytes := response.getHeadersBytes()
	responseBytes = append(responseBytes, headersBytes...)

	// append a new line
	responseBytes = append(responseBytes, []byte("\r\n")...)
	// append body
	responseBytes = append(responseBytes, response.Body...)

	return responseBytes
}

func (response *Response) Send(conn net.Conn) {
	responseBytes := response.getResponseBytes()
	_, err := conn.Write(responseBytes)
	if err != nil {
		fmt.Println("Error sending response: ", err.Error())
		os.Exit(1)
	}
}

func Router(request *Request) Response {
	// default response
	response := Response{
		HttpVersion: request.HttpVersion,
		StatusCode:  Status_NotFound,
		Headers:     map[string]string{},
		Body:        []byte{},
	}

	if request.Path == "/" {
		response.StatusCode = Status_OK
	} else if strings.HasPrefix(request.Path, "/echo/") {
		echoText := strings.TrimPrefix(request.Path, "/echo/")

		fmt.Println("> echo text: ", echoText)
		response.StatusCode = Status_OK
		response.Headers = map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(echoText)),
		}
		response.Body = []byte(echoText)
	} else if request.Path == "/user-agent" {
		userAgent, isUserAgentPresent := request.Headers["user-agent"]
		if isUserAgentPresent {
			// set status code 200
			response.StatusCode = Status_OK
			// set headers
			response.Headers["Content-Type"] = "text/plain"
			response.Headers["Content-Length"] = fmt.Sprint(len(userAgent))
			// set body
			response.Body = []byte(userAgent)
		}
	}

	return response
}

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

	// read request
	requestBytes := make([]byte, 1024)
	requestLength, err := conn.Read(requestBytes)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("> Request Length: ", requestLength)

	requestStr := string(requestBytes[:requestLength])
	fmt.Println(requestStr)

	// parse request
	request := ParseRequest(requestStr)
	fmt.Printf("> parsed request: %+v\n", request)

	// get response
	response := Router(&request)
	fmt.Printf("> generated response: %+v\n", response)

	// send response
	response.Send(conn)
}

package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// HTTPServer 结构体表示一个HTTP服务器
type HTTPServer struct {
	// 监听地址
	addr string
	// 用于等待所有连接处理完成
	wg sync.WaitGroup
}

// NewHTTPServer 创建一个新的HTTP服务器实例
func NewHTTPServer(addr string) *HTTPServer {
	return &HTTPServer{
		addr: addr,
	}
}

// Start 启动HTTP服务器
func (s *HTTPServer) Start() error {
	// 创建TCP监听器
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", s.addr, err)
	}
	defer listener.Close()

	log.Printf("Server started on %s\n", s.addr)

	// 循环接受客户端连接
	for {
		// 接受新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// 为每个连接启动一个新的goroutine处理请求
		s.wg.Add(1)
		go func(c net.Conn) {
			defer s.wg.Done()
			s.handleConnection(c)
		}(conn)
	}
}

// handleConnection 处理单个客户端连接
func (s *HTTPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 使用bufio.Reader读取请求
	reader := bufio.NewReader(conn)

	// 读取第一行（请求行）
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Failed to read request line: %v", err)
		return
	}

	// 解析请求行
	requestLine = strings.TrimSpace(requestLine)
	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		log.Println("Invalid request line")
		return
	}

	method := parts[0]
	path := parts[1]

	// 只处理GET请求
	if method != "GET" {
		s.sendErrorResponse(conn, 405, "Method Not Allowed")
		return
	}

	// 读取请求头（简单处理，只读取到空行）
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading header: %v", err)
			return
		}

		// 遇到空行表示头部结束
		if strings.TrimSpace(line) == "" {
			break
		}
	}

	// 根据路径返回不同的响应
	switch path {
	case "/":
		s.sendResponse(conn, 200, "OK", "text/plain", "Hello, World!")
	case "/health":
		s.sendResponse(conn, 200, "OK", "application/json", `{"status": "healthy"}`)
	default:
		s.sendErrorResponse(conn, 404, "Not Found")
	}
}

// sendResponse 发送成功的HTTP响应
func (s *HTTPServer) sendResponse(conn net.Conn, statusCode int, statusText, contentType, body string) {
	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, statusText)
	response += fmt.Sprintf("Content-Type: %s\r\n", contentType)
	response += fmt.Sprintf("Content-Length: %d\r\n", len(body))
	response += "Connection: close\r\n"
	response += "\r\n"
	response += body

	conn.Write([]byte(response))
}

// sendErrorResponse 发送错误的HTTP响应
func (s *HTTPServer) sendErrorResponse(conn net.Conn, statusCode int, statusText string) {
	body := fmt.Sprintf("<html><body><h1>%d %s</h1></body></html>", statusCode, statusText)
	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, statusText)
	response += "Content-Type: text/html\r\n"
	response += fmt.Sprintf("Content-Length: %d\r\n", len(body))
	response += "Connection: close\r\n"
	response += "\r\n"
	response += body

	conn.Write([]byte(response))
}

package main

import (
	"awesomeProject/internal/server"
	"log"
)

func main() {
	// 创建并启动HTTP服务器
	httpServer := server.NewHTTPServer(":8080")

	// 启动服务器（会一直运行直到手动停止）
	if err := httpServer.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

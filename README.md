# Awesome HTTP Server

一个使用 Go 语言低级套接字编程实现的简单 HTTP 服务器，不依赖 `net/http` 包。

## 功能特性

- ✅ 使用低级 TCP 套接字实现 HTTP 协议
- ✅ 支持并发处理多个客户端连接（Goroutines）
- ✅ 解析基本的 HTTP GET 请求
- ✅ 返回标准 HTTP 响应
- ✅ 支持多路径路由处理
- ✅ 基本错误处理机制

## 项目结构

```
awesomeProject/
├── cmd/
│   └── server/
│       └── main.go              # 程序入口点
├── internal/
│   └── server/
│       └── http_server.go       # HTTP 服务器核心实现
├── go.mod                       # Go 模块定义文件
└── README.md                    # 项目说明文档
```


## 快速开始

### 环境要求

- Go 1.16 或更高版本

### 构建项目

```bash
# 克隆项目
git clone <repository-url>
cd awesomeProject

# 构建可执行文件
go build -o server cmd/server/main.go
```


### 运行服务

```bash
# 方法1: 直接运行
go run cmd/server/main.go

# 方法2: 运行已构建的可执行文件
./server
```


服务器默认在 `localhost:8080` 启动。

## API 接口

- `GET /` - 返回 "Hello, World!" 文本响应
- `GET /health` - 返回 JSON 格式的健康检查状态
- 其他路径 - 返回 404 Not Found 错误页面

### 示例请求

```bash
# 获取主页
curl http://localhost:8080/

# 健康检查
curl http://localhost:8080/health

# 未找到页面
curl http://localhost:8080/notfound
```


## 许可证

MIT License

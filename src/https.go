package main

import (
	"fmt"
	"net"
	"net/http"
)

// -------------------------------------------------------
// ① 用户定义的业务逻辑函数 (最终被调用)
// -------------------------------------------------------
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[YourFunc] 业务逻辑执行中...")
	fmt.Fprintf(w, "Hello astaxie!")
}

// -------------------------------------------------------
// ② main()：注册路由 + 启动 HTTP Server
// -------------------------------------------------------
func main() {
	fmt.Println("[main] 注册路由 '/' -> sayhelloName")
	http.HandleFunc("/", sayhelloName) // 内部包装成 HandlerFunc

	fmt.Println("[main] 启动 HTTP Server，监听 :9090")
	err := http.ListenAndServe(":9090", nil) // nil => DefaultServeMux
	if err != nil {
		panic(err)
	}
}

// -------------------------------------------------------
// 🌍 以下是 Go http 标准库中的执行流程（逻辑示意）
// -------------------------------------------------------

// ListenAndServe 内部逻辑示意：
func ListenAndServe(addr string, handler http.Handler) error {
	fmt.Println("[Server] 创建 Server 实例")
	srv := &http.Server{Addr: addr, Handler: handler}
	return srv.ListenAndServe()
}

// (*Server).ListenAndServe 启动监听：
func (srv *http.Server) ListenAndServe() error {
	fmt.Println("[Server] 开始监听端口", srv.Addr)
	ln, _ := net.Listen("tcp", srv.Addr)
	return srv.Serve(ln)
}

// (*Server).Serve 主循环，等待请求：
func (srv *http.Server) Serve(l net.Listener) error {
	for {
		fmt.Println("[Server] 等待连接...")
		rw, _ := l.Accept() // 阻塞等待 TCP 连接
		fmt.Println("[Server] 收到新连接，开启 goroutine")
		go srv.newConn(rw).serve()
	}
}

// (*conn).serve 处理单个请求：
func (c *conn) serve() {
	fmt.Println("[conn] 解析 HTTP 请求")
	w, r := c.readRequest()

	handler := c.server.Handler // 通常是 DefaultServeMux
	fmt.Println("[conn] 调用 handler.ServeHTTP")
	handler.ServeHTTP(w, r)
}

// (*ServeMux).ServeHTTP 路由分发：
func (mux *http.ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ServeMux] 路由分发开始")
	h, _ := mux.Handler(r) // 查找匹配的 handler
	fmt.Println("[ServeMux] 找到匹配的 Handler，调用 h.ServeHTTP")
	h.ServeHTTP(w, r)
}

// HandlerFunc.ServeHTTP 调用用户函数：
func (f http.HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[HandlerFunc] 调用用户定义函数")
	f(w, r)
}

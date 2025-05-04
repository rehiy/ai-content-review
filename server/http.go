package server

import (
	"log"
	"net/http"
	"os"
)

// 启动 HTTP 服务器
func Start() {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = "0.0.0.0"
	}

	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "8080"
	}

	// 注册路由 /review
	http.HandleFunc("/review", ReviewHandler)

	// 监听指定的地址和端口
	log.Println("启动 HTTP 服务 " + addr + ":" + port)
	err := http.ListenAndServe(addr+":"+port, nil)
	if err != nil {
		panic("启动服务器失败: " + err.Error())
	}
}

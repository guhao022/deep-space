package main

import (
	//"deep-space/utils/log"
	"github.com/num5/env"
)

func main() {
	initialize()

	/*// 获取监听端口
	addr := os.Getenv("LISTEN_ADDR")
	// 监听端口
	listen, err := net.Listen("tcp", addr)
	defer listen.Close()
	if err != nil {
		log.Error("Listen error: %s ", err)
		os.Exit(-1)
	}
	log.Trac("Server start listen on " + addr)

	// 等待客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("Accept error: ", err)
			continue
		}

		// 打印客户端连接
		log.Info(conn.RemoteAddr().String())

		//go connection.NewConnection(conn)
	}*/
}

func initialize() {
	env.Load()
}


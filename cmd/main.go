package main

import (
	"prote-API/pkg/server"
)

var (
	// Listenするアドレス+ポート
	addr string = ":8080"
)

func main() {
	server.Serve(addr)
}

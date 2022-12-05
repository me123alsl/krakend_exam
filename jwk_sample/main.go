package main

import (
	"flag"
	"jwk_sample/application"
)

func main() {

	port := flag.String("port", "8080", "port number")
	flag.Parse()

	// 서버 실행

	webServer := application.Server{Port: *port}
	webServer.RunServer()
}

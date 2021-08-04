package networking

import (
	"fmt"
	"net"
)

func TcpClient() {
	cl, _ := net.Dial("tcp", "127.0.0.1:3000")
	buf := make([]byte, 100)
	for {
		cl.Write([]byte("Hello from TCP client\n"))
		cl.Read(buf)
		fmt.Println(string(buf))
	}
}

func TcpServer() {
	server, _ := net.Listen("tcp", "127.0.0.1:3000")
	buf := make([]byte, 100)
	for {
		cl, _ := server.Accept() //todo: is Accept() blocking the main thread internally
		//handle each connection in their own goroutine
		go func() {
			for {
				cl.Write([]byte("Hello from TCP server\n"))
				cl.Read(buf)
				fmt.Println(string(buf))
			}
		}()
	}
	//TODO: how the main thread is getting blocked
}

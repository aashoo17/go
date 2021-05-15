package networking

import (
	"fmt"
	"net"
)

func TcpClient() {
	cl, _ := net.Dial("tcp", "127.0.0.1:3000")
	buf := make([]uint8, 100)
	for {
		cl.Write([]byte("Hello from TCP client\n"))
		cl.Read(buf)
		fmt.Println(string(buf))
	}
}

func TcpServer() {
	server, _ := net.Listen("tcp", "127.0.0.1:3000")
	buf := make([]uint8, 100)
	for {
		cl, _ := server.Accept()
		//handle each connection in their own goroutine
		go handleConn(cl, buf)
	}
	//TODO: how the main thread is getting blocked
}

func handleConn(cl net.Conn, buf []uint8) {
	for {
		cl.Write([]byte("Hello from TCP client\n"))
		cl.Read(buf)
		fmt.Println(string(buf))
	}
}

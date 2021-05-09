package networking

import (
	"fmt"
	"net"
)

func TcpClient() {
	cl, _ := net.Dial("tcp", "127.0.0.1:3000")
	buf := [100]uint8{}
	for{
		cl.Write([]byte("Hello from TCP client\n"))
		cl.Read(buf[:])
		fmt.Println(buf)
	}
}


func TcpServer(){
	server, _ := net.Listen("tcp", "127.0.0.1:3000")
	buf := [100]uint8{}
	for{
		cl,_ := server.Accept()
		cl.Write([]byte("Hello from TCP client\n"))
		cl.Read(buf[:])
		fmt.Println(buf)
	}
}
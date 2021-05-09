package networking

import (
	"fmt"
	"net"
)

//udp client
func UdpClient() {
	cl, _ := net.Dial("udp", "127.0.0.1:3000")
	buf := [100]uint8{}
	for{
		cl.Write([]byte("Hello from UDP client\n"))
		cl.Read(buf[:])
		fmt.Println(buf)
	}
}

//udp server
func UdpServer(){
	addr,_ := net.ResolveUDPAddr("udp","127.0.0.1:3000")
	server, _ := net.ListenUDP("udp", addr)
	buf := [100]uint8{}
	for{
		server.Write([]byte("Hello from TCP client\n"))
		server.Read(buf[:])
		fmt.Println(buf)
	}
}
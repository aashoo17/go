package networking

import (
	"fmt"
	"net"
)

//udp client
func UdpClient() {
	cl, _ := net.Dial("udp", "127.0.0.1:3000")
	buf := make([]uint8, 100)
	for {
		cl.Write([]byte("Hello from UDP client\n"))
		cl.Read(buf)
		fmt.Println(string(buf))
	}
}

//udp server
func UdpServer() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:3000")
	server, _ := net.ListenUDP("udp", addr)
	buf := make([]uint8, 100)
	//FIXME: fix the server a large goroutine is being spawned over loop and program is crashing
	for {
		go handleUdpConn(server, buf)
	}
}

func handleUdpConn(cl net.Conn, buf []uint8) {
	for {
		cl.Read(buf)
		cl.Write([]byte("Hello from UDP client\n"))
		fmt.Println(string(buf))
	}
}

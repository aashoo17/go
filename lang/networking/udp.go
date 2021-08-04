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
	server, _ := net.ListenUDP("udp", addr) //todo: why this does not take string form of address
	buf := make([]byte, 100)
	ch := make(chan bool)
	go func() {
		for {
			server.Read(buf)
			server.Write([]byte("Hello from UDP client\n"))
			fmt.Println(string(buf))
		}
	}()
	<-ch //block using channel
}

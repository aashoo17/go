package io

import (
	"bufio"
	"fmt"
	"os"
)

func BufWriter() {
	f, _ := os.Create("file.txt")
	//for buffering just create wrapper over File (in general over io.Reader/io.Writer)
	bf := bufio.NewWriter(f)
	defer f.Close()

	bf.Write([]byte("Hello World from buffered IO"))
	bf.Flush()
}

func BufReader() {
	f, _ := os.Open("file.txt")
	bf := bufio.NewReader(f)
	defer f.Close()

	sl := make([]byte, 11)
	bf.Read(sl)

	fmt.Println(string(sl))
}

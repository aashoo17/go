package io

import (
	"bufio"
	"fmt"
	"os"
)

func BufIO() {
	f, _ := os.Create("file.txt")
	//for buffering just create wrapper over File (in general over io.Reader/io.Writer)
	bf := bufio.NewWriter(os.Stdout)
	defer f.Close()

	_, e := bf.Write([]byte("Hello World from buffered IO"))
	//TODO: expected write is not happening - idea is correct though
	//FIXME: fix the write error
	fmt.Println(e)
}

package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

//http get request
func HttpGet() {
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/users")
	buf, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	f, _ := os.Create("resp.txt")
	defer f.Close()
	bufWriter := bufio.NewWriter(f)
	bufWriter.Write(buf)
}

//http post request
func HttpPost() {

}

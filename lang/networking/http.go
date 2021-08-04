package networking

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func Get() {
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	sl, _ := io.ReadAll(resp.Body)

	f, _ := os.Create("file.json")
	writer := bufio.NewWriter(f)
	defer f.Close()
	writer.Write(sl)
}

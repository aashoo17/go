package networking

import (
	"io"
	"net/http"
	"os"
)

func HttpClient() {
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/photos")
	f, _ := os.Create("file.txt")
	sl, _ := io.ReadAll(resp.Body)
	f.Write(sl)
	defer f.Close()
}

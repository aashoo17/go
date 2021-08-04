package main

import (
	"io/fs"
	"os"
)

func Chmod() {
	os.Chmod("resp.txt", fs.ModePerm)
}

func Mkdir() {
	os.Mkdir("dir", fs.ModeDir)
}

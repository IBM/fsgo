package main

import (
	"github.com/IBM/fsgo"
)

func main() {
	var fs = fsgo.NewOsFs()
	b := []byte("bar")
	fs.MkdirAll("some/path/", 0755)
	err := fsgo.WriteFile(fs, "some/path/foo.txt", b, 0644)
	if err != nil {
		panic(err)
	}
	err = fs.Remove("some/path/foo.txt")
	if err != nil {
		panic(err)
	}
	err = fs.RemoveAll("some")
	if err != nil {
		panic(err)
	}
}

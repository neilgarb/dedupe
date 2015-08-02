package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var md5s = make(map[string]string)

func visit(path string, f os.FileInfo, e error) error {
	if f.IsDir() {
		return nil
	}
	var (
		out []byte
		err error
	)
	if runtime.GOOS == "darwin" {
		out, err = exec.Command("md5", "-q", path).Output()
	} else {
		out, err = exec.Command("md5sum", path).Output()
	}
	if err != nil {
		return nil
	}
	md5 := string(out)
	_, ok := md5s[md5]
	if ok {
		fmt.Println(path)
		return nil
	}
	md5s[md5] = path
	return nil
}

func usage() {
	fmt.Println("Usage: dedupe <path>")
}

func main() {
	if len(os.Args) < 2 {
		usage() 
		return
	}
	err := filepath.Walk(os.Args[1], visit)
	if err != nil {
		panic(err)
	}
}

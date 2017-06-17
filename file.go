package main

import (
	"path/filepath"
	"os"
	"fmt"
	"flag"
	"crypto/sha1"
	"io"
)

func getHash(path string) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file error %v\n", err)
		return err
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Printf("io copy error %v\n", err)
		return err
	}
	fmt.Printf("%s hash:%x\n", path, h.Sum(nil))
	return nil
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if (info == nil) {
		return err
	}

	if info.IsDir() {
		return nil
	} else {
		println(path)
		getHash(path)
		return nil
	}
}

func getFilelist(path string) {
	err := filepath.Walk(path, walkFunc)
	if err != nil {
		fmt.Printf("\tfilepath.Walk() return %v\n", err)
	}
}

func main() {
	flag.Parse()
	//root := flag.Arg(0)
	getFilelist("f:\\")
}

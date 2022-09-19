package main

import (
	"flag"
	"fmt"
	sc "main/modules"
	"os"
)

func main() {
	path_config := flag.String("config", "", "specify configuration path")
	flag.Parse()
	if _, err := os.Stat(*path_config); err != nil {
		fmt.Println("invalid path to config")
		return
	}
	fmt.Println("file exists")
	test := []byte("hello")
	fmt.Println(sc.Hash(test))
}

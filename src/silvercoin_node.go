package main

import (
	"flag"
	"fmt"
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

}

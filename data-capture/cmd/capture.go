package main

import (
	"flag"
	"log"
	"unsafe"
)

var configPath = flag.String("f", "", "the path of configuration")

func main() {

	// 启动抓包

	i := 1 << 15
	log.Println(unsafe.Sizeof(i))

}

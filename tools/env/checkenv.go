package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("intのサイズ: %d bytes\n", unsafe.Sizeof(int(0)))
}

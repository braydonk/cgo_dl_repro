package main

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#include "lib.h"
*/
import "C"
import (
	"fmt"
	"log"
)

func main() {
	// handle, err := loadLiblibNow()
	handle, err := loadLiblibLazy()
	if err != nil {
		log.Fatal(err)
	}
	addr, err := dlSym(handle, "get42")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get42 address: ", addr)
	ret := C.get42()
	fmt.Println(ret)
}

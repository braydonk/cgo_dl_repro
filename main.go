package main

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#include "lib.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	// handle, err := loadLiblibNow()
	// handle, err := loadLiblibLazy()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// get42Addr, err := dlSym(handle, "get42")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("get42 address: ", get42Addr)
	ret := C.get42()
	fmt.Println(ret)
}

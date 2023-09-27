package main

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#include "lib.h"
*/
import "C"
import "fmt"

func main() {
	loadLiblib()
	ret := C.get42()
	fmt.Println(ret)
}

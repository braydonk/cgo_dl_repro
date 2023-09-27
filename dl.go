package main

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func loadLiblibLazy() (unsafe.Pointer, error) {
	return loadLiblib(C.RTLD_LAZY)
}

func loadLiblibNow() (unsafe.Pointer, error) {
	return loadLiblib(C.RTLD_NOW)
}

func loadLiblib(nowOrLazy int) (unsafe.Pointer, error) {
	name := C.CString("./liblib.so")
	defer C.free(unsafe.Pointer(name))
	handle := C.dlopen(name, C.int(nowOrLazy|C.RTLD_GLOBAL))
	if handle == nil {
		return nil, dlError()
	}
	return handle, nil
}

func dlError() error {
	lastErr := C.dlerror()
	if lastErr == nil {
		return nil
	}
	return errors.New(C.GoString(lastErr))
}

func dlSym(handle unsafe.Pointer, symbol string) (unsafe.Pointer, error) {
	sym := C.CString(symbol)
	defer C.free(unsafe.Pointer(sym))
	dlError()
	pointer := C.dlsym(handle, sym)
	if pointer == nil {
		return nil, fmt.Errorf("symbol %q not found: %w", symbol, dlError())
	}
	return pointer, nil
}

package main

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

func loadLiblib() (unsafe.Pointer, error) {
	name := C.CString("./liblib.so")
	defer C.free(unsafe.Pointer(name))
	var handle unsafe.Pointer

	if err := withOSLock(func() error {
		handle = C.dlopen(name, C.int(C.RTLD_LAZY|C.RTLD_GLOBAL))
		if handle == nil {
			return dlError()
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return handle, nil
}

func withOSLock(action func() error) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	return action()
}

func dlError() error {
	lastErr := C.dlerror()
	if lastErr == nil {
		return nil
	}
	return errors.New(C.GoString(lastErr))
}

// +build go1.7,!go1.8

package gid

import "unsafe"

// In order to make easy life in asm, we need to copy the struct from golang
// source code https://github.com/golang/go/blob/release-branch.go1.7/src/runtime/runtime2.go

type g struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr

	_panic       uintptr
	_defer       uintptr
	m            uintptr
	stackAlloc   uintptr
	sched        gobuf
	syscallsp    uintptr
	syscallpc    uintptr
	stkbar       []stkbar
	stkbarPos    uintptr
	stktopsp     uintptr
	param        unsafe.Pointer
	atomicstatus uint32
	stackLock    uint32
	goid         int64 // <= go!
}

type stack struct {
	lo uintptr
	hi uintptr
}

type stkbar struct {
	savedLRPtr uintptr
	savedLRVal uintptr
}

type gobuf struct {
	sp   uintptr
	pc   uintptr
	g    uintptr
	ctxt unsafe.Pointer
	ret  uint64
	lr   uintptr
	bp   uintptr
}

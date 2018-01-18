package gid

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
)

const goroutine = "goroutine"

func isdigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// gidFromStack restrieves goroutine id from stack information, like:
//
//	goroutine xxx [running]:
//		....
//
func gidFromStack() int64 {
	buf := make([]byte, 100)
	runtime.Stack(buf, false)

	buf = buf[len(goroutine)+1:]

	idx := 0
	for idx < len(buf) && isdigit(buf[idx]) {
		idx += 1
	}

	id, err := strconv.ParseInt(string(buf[:idx]), 10, 64)
	if err != nil {
		panic(err)
	}
	return id
}

func TestGoID(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			gotId := GoID()
			expectedId := gidFromStack()
			if gotId != expectedId {
				t.Errorf("expectedID %d, but got %d\n", expectedId, gotId)
			}
		}()
	}

	wg.Wait()
}

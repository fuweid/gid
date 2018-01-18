// +build darwin linux
// +build amd64

package gid

// getgid restrieves goroutine id from TSL.
func getgid() int64

// GoID gets current goroutine ID.
func GoID() int64 {
	return getgid()
}

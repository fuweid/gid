## gid for fun

gid is used to get current goroutine id from os-thread.

### get started

This package has only one exported function, `GoID`. Easy to remember, right? 😀

```
func doXXX() {
        id := gid.GoID() // get current goroutine id
	      ....
}
```

### Support?

- [x] go 1.7.x @ amd64 linux
- [x] go 1.8.x @ amd64 linux
- [x] go 1.9.x @ amd64 linux

> Plan to support osx... but travis osx bootup is too slow for testing 😂. I will add it when testing env is ready.

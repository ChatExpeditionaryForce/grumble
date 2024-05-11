package main

import "unsafe"

// Stolen from https://dev.to/chigbeef_77/bool-int-but-stupid-in-go-3jb3
func b2i(b bool) int {
	return int(*(*byte)(unsafe.Pointer(&b)))
}

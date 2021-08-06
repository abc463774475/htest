package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	cStr := C.CString("Hello, 世界\n")
	defer C.free(unsafe.Pointer(cStr))
	C.puts(C.CString("Hello, 世界\n"))
}
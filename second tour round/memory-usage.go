package main

import (
    "fmt"
    "unsafe"
)

// Memory usage. For example, types struct{}, interface{}, and bool,
func main() {
    var s struct{}
    fmt.Println(unsafe.Sizeof(s))
    var i interface{}
    fmt.Println(unsafe.Sizeof(i))
    var b bool
    fmt.Println(unsafe.Sizeof(b))
}
package main

import (
	"fmt"
)

type Temp struct {
}

func main() {
	var pnt *Temp       // pointer
	var inf interface{} // interface declaration
	fmt.Printf("1 inf is a nil interface of type %T : %v\n", inf, inf == nil)
	inf = nil
	fmt.Printf("2 inf is a nil interface of type %T : %v\n", inf, inf == nil)
	inf = pnt // inf is a non-nil interface holding a nil pointer (pnt)

	fmt.Printf("pnt is a nil pointer: %v\n", pnt == nil)
	fmt.Printf("3 inf is a nil interface of type %T : %v\n", inf, inf == nil)
	fmt.Printf("inf is a interface of type %T holding a nil pointer: %v\n", inf, inf == (*Temp)(nil))
}

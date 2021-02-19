package main

import "fmt"

func main(){
	// type dclaratiaon adalah type alian untuk mendelegasikan satu type ke suatu identifier
	// dengan tambahan (type)

	type isMarried bool

	var afi isMarried = false
	fmt.Println(afi)
}
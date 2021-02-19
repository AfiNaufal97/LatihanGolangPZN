package main

import "fmt"

func main(){
	var nilai1 int8 = 80
	var nilai2 int8 = 80
	var nilai3 int8 = 90

	// operasi boolean && || !
	fmt.Println((nilai1 == nilai2) && (nilai2 < nilai3))
	fmt.Println((nilai1 <= nilai2) || (nilai2 == nilai3))
	fmt.Println(!((nilai1 > nilai2) && (nilai2 == nilai3)))
}
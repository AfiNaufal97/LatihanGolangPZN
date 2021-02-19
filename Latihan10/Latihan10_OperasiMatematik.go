package main

import "fmt"
func main(){
	var nilai1 int = 10
	var nilai2 int = 10

	// hitung
	var hasil int = nilai1 + nilai2
	fmt.Println(hasil)

	// dengan operasi unary(+=  ++   --   !)
	// opersi yang hanya memerlukan dirinya sendiri
	hasil++
	fmt.Println(hasil)
	hasil--
	fmt.Println(hasil)

	hasil += 10
	fmt.Println(hasil)
}
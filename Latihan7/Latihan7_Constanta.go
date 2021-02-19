package main

import "fmt"

func main(){
	// dengan constanta kita tidak wajib memanggil variable yang constanta
	// beberda dengan variable non konstanta yang wajib di panggil
	// konstanta tidak dapat di rubah
	const NAMA string = "AFI NAUFAL SETIAWAN"
	fmt.Println(NAMA)

	// NAMA = "AFI" -> error karena constanta tidak bisa di rubah
	fmt.Println(NAMA)
}
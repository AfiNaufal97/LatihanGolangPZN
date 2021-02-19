package main

import "fmt"

func main(){
	// dalam golang variable ada dalam var(variable)
	var nama string = "Afi Naufal Setiawan"
	fmt.Println(nama)

	// varibale dalam penggunaanya ketiak di deklarasikan wajib di panggil, jika tidak akan merah pada source code

	// dalam pembuatan variable bisa juga tanpa var
	// tetapi menggunakan (:=)
	nama2 := "Afi"
	fmt.Println(nama2)
}
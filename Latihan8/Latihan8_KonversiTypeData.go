package main

import "fmt"

func main(){
	// konversi type data bisa dilakukan asalak dat ayng di konversi dari 1 type ke type lain itu sesuai

	// misal
	var nilai int8 = 100

	// casting / konversikan ke int
	var nilaiString string = string(nilai)

	// cetak
	fmt.Println(nilaiString)
}
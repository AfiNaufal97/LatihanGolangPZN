package main

import "fmt"

func fungsiUtama(masukan string, masukanKedua func(nilai string)string)string{
	hasil := masukanKedua(masukan)
	return hasil
}


// ini seperti type alias = untuk mempersingkat code yang panjang
// asalkan kode yang si singkat itu relevan dengan type alias
type Masukan2 func(input string)string

func fungsiUtama2(masukan string, masukan2 Masukan2)string{
	hasil := masukan2(masukan)
	return hasil
}

func kedua(masukan string)string{
	return masukan
}

func main(){

	// eksekusi
	hasilAkhir := fungsiUtama("Afi", kedua)
	fmt.Println(hasilAkhir)

	hasilKedua := fungsiUtama2("Afi", kedua)
	fmt.Println(hasilKedua)
}
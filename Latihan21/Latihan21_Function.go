package main

import "fmt"

// function dalam golang di awali dengan func
// dan nama sebaiknya camelcase contoh : (kalimatPertama)
// dalam satu modul.go hanya boleh ada 1 function main, karena funt main adalah funtion yang akan menjalankan program
func functionSatu(){
	fmt.Println("Dari Afi")
}

func main(){
	// eksekusi function
	functionSatu()

	// bisa juga dengan looping
	for i:=0; i<10; i++{
		functionSatu()
	}
}
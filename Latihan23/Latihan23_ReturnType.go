package main

import "fmt"

// membuat function dengan retun type di golang
func nilaiHasil(masukan int)int{
	return masukan*2
}

// retur dengan nilai balikan lebih dari 1
func banyakBalikan()(int, int, int){
	return 10, 10, 10
}

func main(){
	nilai := nilaiHasil(10)
	fmt.Println(nilai)

	satu, dua, tiga := banyakBalikan()
	fmt.Println(satu)
	fmt.Println(dua)
	fmt.Println(tiga)

}
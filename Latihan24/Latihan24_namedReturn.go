package main

import "fmt"

// freturn type dengan nama
func banyakData()(balikan1 string, balikan2 string, balikan3 string){
	balikan1 = "satu"
	balikan2 = "dua"
	balikan3 = "tiga"
	return
}

func main(){
	// eksekusi semua
	nilai1, nilai2, nilai3 := banyakData()
	fmt.Println(nilai1) 
	fmt.Println(nilai2) 
	fmt.Println(nilai3) 

	// eksekusi bebrapa dsta saja
	nilai4, _, nilai6 := banyakData()
	fmt.Println(nilai4) 
	fmt.Println(nilai6) 
}
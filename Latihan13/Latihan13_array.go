package main

import "fmt"

func main(){
	// array adalah sebuah type data yang dapat menampung banyak data
	var nilai = [5]int{1, 2, 3, 4, 5}

	fmt.Println(nilai)
	fmt.Println(nilai[3])
	fmt.Println(len(nilai))

	var nilai2 [3]int
	nilai2[0] = 1
	nilai2[1] = 1
	nilai2[2] = 1
	fmt.Println(nilai2)
}
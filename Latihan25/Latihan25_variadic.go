package main

import "fmt"

// function yang dapat menampung banyak data tetpi data buka berupa array atau slice
// disebut variadic
// diletakan paling akhir jika ada paremeter lainnya
func variadicFun(latihan... int)[]int{
	return latihan
}

func main(){
	dataSementara := variadicFun(10, 20, 30)
	fmt.Println(dataSementara)

	// jikanilai masukan berupa slice juga bisa, tetpi di tambahkan tanda (...)
	dataSlice := []int{1, 2, 3, 4, 5}
	dataIsi := variadicFun(dataSlice ...)
	fmt.Println(dataIsi) 
}
package main

import "fmt"

func main(){
	// slice adalah tipe dta yang merupakan potongan dari array
	// jika di array hanya memiliki method length saja
	// maka di slice memilikimethod pointer, length dan copacity

	// cara membuta slice

	// 1. buat array dahulu
	var array1= [...]int{
		1,
		2,
		3,
	}


	// 2. membuat slice dengan menganbil bagian dari array

	var sliceSatu = array1
	fmt.Println(sliceSatu)
	fmt.Println(sliceSatu[:])

	// ambil panjang slice
	fmt.Println(len(sliceSatu))

	// ambil capacity
	fmt.Println(cap(sliceSatu))

	// ambil pointer
	fmt.Println(sliceSatu[0])


	// tambah slice
	//  buta slice dahulu
	 bulanSatu := [...]string{
		"januari",
		"februari",
		"maret", 
		"april", 
		"mei", 
		"juni",
		"juli",
		"agustus",
		"september", 
		"oktober", 
		"november",
		"desember",
	}

	var sliceBulan = bulanSatu[10:]
	fmt.Println(sliceBulan)

	// add data dengan append
	sliceBulan.append(bulanSatu, )
}
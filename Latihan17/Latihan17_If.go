package main

import "fmt"

func main(){
	// latihan if statment 

	var nama string = "Afi"
	if nama == "Afi"{ // jika hasil perbadingan tru maka akan dieksekusi, jika tidak maka akan diabaikan(tidak dieksekusi)
		fmt.Println("hallo "+nama)
	}


	// if statment dengan short deklarasi
	if nama2 := len(nama); nama2 > 5{
		fmt.Println(nama+" terlalu panjang")
	}


	// if dengan else
	if nama == "Afi"{ // jika hasil perbadingan tru maka akan dieksekusi, jika tidak maka akan diabaikan(tidak dieksekusi)
		fmt.Println("hallo mas "+nama)
	}else if nama == "afi"{
		fmt.Println("hallo mas "+nama)
	}else{
		fmt.Println("hallo "+nama)
	}

	// if statment dengan short deklarasi dengan else
	if nama2 := len(nama); nama2 > 5{
		fmt.Println(nama+" terlalu panjang")
	}else if nama2 > 3{
		fmt.Println(nama+" cukup panjang")
	}else{
		fmt.Println(nama+" cukup baik")
	}

}
package main

import "fmt"

// buat function dengan paremeter di golang dengan nilai balikan
func nilaiMasukan(nilai int)bool{
	if(nilai == 90){
		return true
	}else{
		return false
	}
}

func main(){
	hasil := nilaiMasukan(10)
	fmt.Println(hasil)
}
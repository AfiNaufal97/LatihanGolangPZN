package main

import "fmt"

func main(){
	// break atau menghentikan paksa kode
	for i:=0; i< 10; i++{
		// kode akan dihentikan paksa jika nilai i == 5
		if(i == 5){
			break
		}
		fmt.Println(i)
	}

	// continue nilai( akan di skip, ke nilai selanjutnya)
	for j:=0 ; j < 10; j++{
		if(j % 2 == 0){
			continue
		}
		fmt.Println(j)
	}
}
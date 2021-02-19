package main

import "fmt"

// swict adalah sebuah if dengan ar lebih sederhana
func main(){

	var nama string = "Afi"
	switch nama {
	case "Afi":
		fmt.Println("hallo "+nama)
	case "afi":
		fmt.Println("hallo "+nama)
	default:
		fmt.Println("hallo ")
	}

	// switch case dengan short deklarasi
	switch nama2:= len(nama); nama2 > 5 {
	case true:
		fmt.Println("hallo "+nama)
	case false:
		fmt.Println("hallo ")
	}

	// switc case tanpa identifier
	var nama3 int = len(nama)
	switch{
	case nama3 == 3:
		fmt.Println("Hallo "+nama)
	case nama3 == 5:
		fmt.Println("Hallo")
	}

	
}
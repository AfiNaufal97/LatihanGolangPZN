package main

import "fmt"

func main(){

	// deklarasi map
	// mapa adalah kumpulan key dan value
	// var identifier map[type data key]type data value

	var keterangan map[string]string = map[string]string{
		"nama" : "Afi Naufal Setiawan",
		"kelas" : "5A",
		"nim" : "20092001",
	}
	fmt.Println(keterangan)
	fmt.Println(len(keterangan))

	// tambah data
	// jika tidak ada keyyang sama maka akan ditambah
	// tetapi jika key sama,maka key tersebut isiny akan di replace dengan key yang terakhir kali dimasukan
	keterangan["alamat"] = "Tegal"
	fmt.Println(keterangan)
	fmt.Println(len(keterangan))


	// map dengan make
	hasil := make(map[string]string)
	hasil["kesatu"] = "Satu"
	hasil["kedua"] = "Dua"
	hasil["ketiga"] = "Tiga"
	fmt.Println(hasil)
	fmt.Println(len(hasil))

	// delete
	delete(hasil, "kesatu")
	fmt.Println(hasil)



}
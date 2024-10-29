package main

import "fmt"

func main() {
	var mk string = "Algoritma dan Pemrograman"
	var sks int = 3
	var kode string = "CAK123"
	fmt.Println("Tuliskan kode MK dan SKS")
	fmt.Scan(&kode, &sks)
	fmt.Println("Kredit MK", kode, "-", mk, "1 adalah", sks, "SKS")
}

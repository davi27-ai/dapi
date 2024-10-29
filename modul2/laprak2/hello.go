package main

import "fmt"

func main() {
	var (
		nama  string
		prodi string
		kelas string
		nim   string
	)

	fmt.Println("masukan nama")
	fmt.Scan(&nama)

	fmt.Println("masukan kelas")
	fmt.Scan(&kelas)

	fmt.Println("masukan nim")
	fmt.Scan(&nim)

	fmt.Println("Perkenalkan saya adalah", nama, "salah satu mahasiswa prodi S1 teknik informatika", prodi,
		"dari kelas", kelas, "dengan NIM", nim)
}
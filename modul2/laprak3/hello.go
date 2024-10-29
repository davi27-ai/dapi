package main

import "fmt"

func main() {
	var huruf int
	var konsonan bool

	fmt.Scanf("%c", &huruf)
	konsonan = huruf != 'b' && huruf != 'c' && huruf != 'd' && huruf != 'f' && huruf != 'g'
	fmt.Println(konsonan)
}
package main

import "fmt"

func main() {
	var cel, fah int64

	fmt.Print("masukan fahrenheit = ")
	fmt.Scanln(&fah)

	cel = (fah - 32) * 5 / 9

	fmt.Print(cel)
}

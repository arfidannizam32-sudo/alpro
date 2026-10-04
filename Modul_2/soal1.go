package main

import "fmt"

func main() {
	var nama, kelas string
	var nim int64

	fmt.Print("input nama=")
	fmt.Scanln(&nama)

	fmt.Print("input nim=")
	fmt.Scanln(&nim)

	fmt.Print("input kelas=")
	fmt.Scanln(&kelas)

	fmt.Printf("perkenalan nama saya %s, salah satu mahasiswa S1-IF dari kelas %s dengan NIM %d.", nama, kelas, nim)
}

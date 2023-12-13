package main

import (
	"fmt"
	"modul/read"
	"modul/task"
)

func main() {
	info, err := read.OpenJson()
	if err != nil {
		panic(err)
	}

	// Barcha customerlarning umumiy balansi va umumiy harajjatlarini chiqarish kerak

	fmt.Print("1) ")
	task.Task1(info)
	fmt.Println()

	// Eng ko'p pul sarf etganni topish

	fmt.Print("2) ")
	task.Task2(info)
	fmt.Println()

	// Eng qimmat mahsulotni aniqlash

	fmt.Print("3) ")
	task.Task3(info)
	fmt.Println()

	// Barcha productlarning urtacha narxi

	fmt.Print("4) ")
	task.Task4(info)
	fmt.Println()

	// Eng arzon savdo qilgan customer

	fmt.Print("5) ")
	task.Task5(info)
	fmt.Println()

	// Eng ko'p sotilgan productlar kategiriyasi

	fmt.Print("6) ")
	task.Task6(info)
	fmt.Println()

	// Eng ko'p sotilgan mahsulotni topish

	fmt.Print("7) ")
	task.Task7(info)
	fmt.Println()

	// Barcha mahsulotlarning o'rtacha miqdori

	fmt.Print("8) ")
	task.Task8(info)
	fmt.Println()

	// Eng ko'p mahsulot sotib olgan mijoz

	fmt.Print("9) ")
	task.Task9(info)
	fmt.Println()

	// Eng ko'p savdolarda ko'rinadigan mahsulot

	fmt.Print("10) ")
	task.Task10(info)
	fmt.Println()

	// O'rtacha savdo mablag'i

	fmt.Print("11) ")
	task.Task11(info)
	fmt.Println()

	// Eng ko'p umumiy daromad qilgan toifani aniqlash

	fmt.Print("12) ")
	task.Task12(info)
	fmt.Println()

	// Eng qimmat productni o'z ichiga olgan savdo

	fmt.Print("13) ")
	task.Task13(info)
	fmt.Println()

	fmt.Print("14) ")
	task.Task14(info)
	fmt.Println()

	// Umumiy nechta sotilganligi

	fmt.Print("15) ")
	task.Task15(info)

}

package main

import (
	"github.com/NaufalWaiz/calculator-package"
	"fmt"
)

func main() {
	Penjumlahan()
	Pengurangan()
	Perkalian()
	Pembagian()
}

func Penjumlahan() {	
	calculator.Penjumlahan(1, 2)
	fmt.Println("Hasil Penjumlahan : ", calculator.Penjumlahan(1, 2))
}

func Pengurangan() {
	calculator.Pengurangan(2, 1)
	fmt.Println("Hasil Pengurangan : ", calculator.Pengurangan(2, 1))
}

func Perkalian() {
	calculator.Perkalian(2, 2)
	fmt.Println("Hasil Perkalian : ", calculator.Perkalian(2, 2))
}

func Pembagian() {
	calculator.Pembagian(4, 2)
	result, err := calculator.Pembagian(4, 2)
	if err != nil {
		fmt.Println("Error Pembagian : ", err)
	} else {
		fmt.Println("Hasil Pembagian : ", result)
	}
}
package main

import (
	"fmt"
)
func main() {
	// Closure untuk menentukan nilai huruf
	getNilaiHuruf := func(nilaiUjian int) string {
		batasNilai := map[string]int{
			"A": 90,
			"B": 80,
			"C": 70,
			"D": 60,
		}
		for huruf, batas := range batasNilai {
			if nilaiUjian >= batas {
				return huruf
			}
		}
		return "E"
	}

	// Menentukan nilai huruf untuk beberapa nilai ujian
	nilaiUjian := []int{85, 75, 95, 55, 65}
	for _, nilai := range nilaiUjian {
		fmt.Println("Nilai ujian", nilai, ":", getNilaiHuruf(nilai))
	}
}
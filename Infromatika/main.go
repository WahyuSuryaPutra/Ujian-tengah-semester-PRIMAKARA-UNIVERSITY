package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x int
	if len(os.Args) < 2 {
		fmt.Println("Silahkan masukan argument si X")
	} else {
		inputX, convertToIntError := strconv.Atoi(os.Args[1])
		if convertToIntError != nil {
			fmt.Println("Input yang dimasukan bukan merupakan angka")
			os.Exit(1)
		}
		x = inputX

		i := 0
		a := 7
		b := 9
		c := 100

		f := c % 3

		for i < 100 {
			x = x - 20
			b = b - 2
			if b > 0 {
				b = b + 1
				a = a + 1
				f = f + a
			} else {
				a = a + 1
				f = f + a
			}
			for f%4 == 0 {
				b = b + 1
				a = a + 1
				f = f + a
			}
			i = i + 10
		}

		b = b * 2

		fmt.Printf("Nilai A:%d\nNilai B:%d\nNilai C:%d\nNilai F:%d\nNilai X:%d", a, b, c, f, x)

		os.Exit(0)
	}
}

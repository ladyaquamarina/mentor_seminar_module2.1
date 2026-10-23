package main

import (
	"fmt"
	"os"

	"go_module1/console"
)

func printTable(n int) {
	fmt.Printf("Таблица умножения для %d:\n", n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %2d = %d\n", n, i, n*i)
	}
}

func main() {
	n, err := console.New(os.Stdin, os.Stdout).Int("Введите число: ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
	printTable(n)
}

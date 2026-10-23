package main

import (
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for d := 3; d*d <= n; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func run(in io.Reader, out io.Writer) error {
	n, err := console.New(in, out).Int("Введите целое число: ")
	if err != nil {
		return err
	}
	if isPrime(n) {
		fmt.Fprintf(out, "%d — простое число\n", n)
	} else {
		fmt.Fprintf(out, "%d — не является простым числом\n", n)
	}
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

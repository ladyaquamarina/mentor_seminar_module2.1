package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

var ErrInvalidN = errors.New("N должно быть натуральным числом (N ≥ 1)")

func SumEvenDivBy3(n int) (sum int, matched []int, err error) {
	if n < 1 {
		return 0, nil, ErrInvalidN
	}
	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
			matched = append(matched, i)
		}
	}
	return sum, matched, nil
}

func run(in io.Reader, out io.Writer) error {
	n, err := console.New(in, out).Int("Введите N: ")
	if err != nil {
		return err
	}
	sum, matched, err := SumEvenDivBy3(n)
	if err != nil {
		return err
	}
	if len(matched) == 0 {
		fmt.Fprintf(out, "В диапазоне от 1 до %d нет чисел, которые чётные и делятся на 3\n", n)
	} else {
		fmt.Fprintf(out, "Подходящие числа: %v\n", matched)
	}
	fmt.Fprintf(out, "Сумма: %d\n", sum)
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

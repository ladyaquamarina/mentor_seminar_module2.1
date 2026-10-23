package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

var ErrInvalidN = errors.New("N должно быть натуральным числом (N ≥ 1)")

func SumAndAverage(n int) (sum int, avg float64, err error) {
	if n < 1 {
		return 0, 0, ErrInvalidN
	}
	for i := 1; i <= n; i++ {
		sum += i
	}
	avg = float64(sum) / float64(n)
	return sum, avg, nil
}

func run(in io.Reader, out io.Writer) error {
	n, err := console.New(in, out).Int("Введите N: ")
	if err != nil {
		return err
	}
	sum, avg, err := SumAndAverage(n)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Сумма чисел от 1 до %d: %d\n", n, sum)
	fmt.Fprintf(out, "Среднее значение: %s\n", console.FormatNumber(avg))
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

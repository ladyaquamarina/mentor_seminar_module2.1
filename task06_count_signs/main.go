package main

import (
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

type Counts struct {
	Positive int
	Negative int
	Zero     int
}

func CountSigns(nums []float64) Counts {
	var c Counts
	for _, v := range nums {
		if v > 0 {
			c.Positive++
		} else if v < 0 {
			c.Negative++
		} else {
			c.Zero++
		}
	}
	return c
}

func run(in io.Reader, out io.Writer) error {
	r := console.New(in, out)
	n, err := r.Count("Введите количество чисел N: ")
	if err != nil {
		return err
	}
	nums, err := r.Floats(n)
	if err != nil {
		return err
	}
	c := CountSigns(nums)
	fmt.Fprintf(out, "Положительных: %d\n", c.Positive)
	fmt.Fprintf(out, "Отрицательных: %d\n", c.Negative)
	fmt.Fprintf(out, "Нулей: %d\n", c.Zero)
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

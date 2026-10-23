package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

var ErrEmpty = errors.New("нет чисел для сравнения")

func MinMax(nums []float64) (min, max float64, err error) {
	if len(nums) == 0 {
		return 0, 0, ErrEmpty
	}

	min, max = nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max, nil
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
	min, max, err := MinMax(nums)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Максимальное значение: %s\n", console.FormatNumber(max))
	fmt.Fprintf(out, "Минимальное значение: %s\n", console.FormatNumber(min))
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

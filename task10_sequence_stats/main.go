package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

var ErrEmpty = errors.New("последовательность пуста")

type Stats struct {
	Count   int
	Sum     int
	Average float64
	Max     int
	Min     int
	Even    int
	Odd     int
}

func Calculate(nums []int) (Stats, error) {
	if len(nums) == 0 {
		return Stats{}, ErrEmpty
	}
	s := Stats{Count: len(nums), Max: nums[0], Min: nums[0]}
	for _, v := range nums {
		s.Sum += v
		if v > s.Max {
			s.Max = v
		}
		if v < s.Min {
			s.Min = v
		}
		if v%2 == 0 {
			s.Even++
		} else {
			s.Odd++
		}
	}
	s.Average = float64(s.Sum) / float64(s.Count)
	return s, nil
}

func run(in io.Reader, out io.Writer) error {
	r := console.New(in, out)
	n, err := r.Count("Введите количество чисел N: ")
	if err != nil {
		return err
	}
	nums, err := r.Ints(n)
	if err != nil {
		return err
	}
	s, err := Calculate(nums)
	if err != nil {
		return err
	}

	fmt.Fprintln(out, "Статистика последовательности:")
	fmt.Fprintf(out, "  сумма: %d\n", s.Sum)
	fmt.Fprintf(out, "  среднее: %s\n", console.FormatNumber(s.Average))
	fmt.Fprintf(out, "  максимум: %d\n", s.Max)
	fmt.Fprintf(out, "  минимум: %d\n", s.Min)
	fmt.Fprintf(out, "  чётных: %d\n", s.Even)
	fmt.Fprintf(out, "  нечётных: %d\n", s.Odd)
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

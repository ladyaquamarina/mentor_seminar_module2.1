package main

import (
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

type Analysis struct {
	Sign       string
	Even       bool
	DivBy3     bool
	DivBy5     bool
	DivBy3And5 bool
}

func Analyze(n int) Analysis {
	var a Analysis

	if n > 0 {
		a.Sign = "положительное"
	} else if n < 0 {
		a.Sign = "отрицательное"
	} else {
		a.Sign = "ноль (ни положительное, ни отрицательное)"
	}

	a.Even = n%2 == 0
	a.DivBy3 = n%3 == 0
	a.DivBy5 = n%5 == 0
	a.DivBy3And5 = a.DivBy3 && a.DivBy5
	return a
}

func yesNo(b bool) string {
	if b {
		return "да"
	}
	return "нет"
}

func run(in io.Reader, out io.Writer) error {
	n, err := console.New(in, out).Int("Введите целое число N: ")
	if err != nil {
		return err
	}
	a := Analyze(n)

	parity := "нечётное"
	if a.Even {
		parity = "чётное"
	}

	fmt.Fprintf(out, "Число %d:\n", n)
	fmt.Fprintf(out, "  знак: %s\n", a.Sign)
	fmt.Fprintf(out, "  чётность: %s\n", parity)
	fmt.Fprintf(out, "  делится на 3: %s\n", yesNo(a.DivBy3))
	fmt.Fprintf(out, "  делится на 5: %s\n", yesNo(a.DivBy5))
	fmt.Fprintf(out, "  делится и на 3, и на 5: %s\n", yesNo(a.DivBy3And5))
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

var (
	ErrDivisionByZero  = errors.New("деление на ноль невозможно")
	ErrUnknownOperator = errors.New("неизвестная операция, допустимы: + - * /")
)

func Calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrUnknownOperator
	}
}

func run(in io.Reader, out io.Writer) error {
	r := console.New(in, out)

	a, err := r.Float("Введите первое число: ")
	if err != nil {
		return err
	}
	op, err := r.Word("Введите операцию (+, -, *, /): ")
	if err != nil {
		return err
	}
	b, err := r.Float("Введите второе число: ")
	if err != nil {
		return err
	}

	result, err := Calculate(a, b, op)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Результат: %s %s %s = %s\n",
		console.FormatNumber(a), op, console.FormatNumber(b), console.FormatNumber(result))
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

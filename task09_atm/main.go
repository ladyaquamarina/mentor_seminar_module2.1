package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"go_module1/console"
)

var (
	ErrInvalidAmount     = errors.New("сумма должна быть больше нуля")
	ErrInsufficientFunds = errors.New("недостаточно средств на балансе")
	ErrNegativeBalance   = errors.New("начальный баланс не может быть отрицательным")
)

type Account struct {
	Balance float64
}

func NewAccount(initial float64) (*Account, error) {
	if initial < 0 {
		return nil, ErrNegativeBalance
	}
	return &Account{Balance: initial}, nil
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > a.Balance {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	return nil
}

const menu = `
Выберите действие:
  1 — посмотреть баланс
  2 — пополнить баланс
  3 — снять средства
  0 — выход
> `

func money(v float64) string {
	return fmt.Sprintf("%.2f", v)
}

func run(in io.Reader, out io.Writer) error {
	r := console.New(in, out)

	initial, err := r.Float("Введите начальный баланс: ")
	if err != nil {
		return err
	}
	acc, err := NewAccount(initial)
	if err != nil {
		return err
	}

	for {
		choice, err := r.Int(menu)
		if err != nil {
			return err
		}

		switch choice {
		case 1:
			fmt.Fprintf(out, "Текущий баланс: %s\n", money(acc.Balance))
		case 2:
			amount, err := r.Float("Сумма пополнения: ")
			if err != nil {
				return err
			}
			if err := acc.Deposit(amount); err != nil {
				fmt.Fprintln(out, "Ошибка:", err)
			} else {
				fmt.Fprintf(out, "Баланс пополнен. Текущий баланс: %s\n", money(acc.Balance))
			}
		case 3:
			amount, err := r.Float("Сумма снятия: ")
			if err != nil {
				return err
			}
			if err := acc.Withdraw(amount); err != nil {
				fmt.Fprintf(out, "Ошибка: %v (доступно %s)\n", err, money(acc.Balance))
			} else {
				fmt.Fprintf(out, "Выдано %s. Текущий баланс: %s\n", money(amount), money(acc.Balance))
			}
		case 0:
			fmt.Fprintln(out, "До свидания!")
			return nil
		default:
			fmt.Fprintln(out, "Неизвестное действие, выберите 0–3")
		}
	}
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "\nОшибка:", err)
		os.Exit(1)
	}
}

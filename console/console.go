package console

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
)

type Reader struct {
	in  *bufio.Reader
	out io.Writer
}

func New(in io.Reader, out io.Writer) *Reader {
	return &Reader{in: bufio.NewReader(in), out: out}
}

func (r *Reader) Int(prompt string) (int, error) {
	fmt.Fprint(r.out, prompt)
	var v int
	if _, err := fmt.Fscan(r.in, &v); err != nil {
		return 0, wrap("ожидалось целое число", err)
	}
	return v, nil
}

func (r *Reader) Float(prompt string) (float64, error) {
	fmt.Fprint(r.out, prompt)
	var v float64
	if _, err := fmt.Fscan(r.in, &v); err != nil {
		return 0, wrap("ожидалось число", err)
	}
	return v, nil
}

func (r *Reader) Word(prompt string) (string, error) {
	fmt.Fprint(r.out, prompt)
	var v string
	if _, err := fmt.Fscan(r.in, &v); err != nil {
		return "", wrap("ожидалось значение", err)
	}
	return v, nil
}

func (r *Reader) Count(prompt string) (int, error) {
	n, err := r.Int(prompt)
	if err != nil {
		return 0, err
	}
	if n <= 0 {
		return 0, fmt.Errorf("количество должно быть больше нуля, получено %d", n)
	}
	return n, nil
}

func (r *Reader) Floats(n int) ([]float64, error) {
	nums := make([]float64, 0, n)
	for i := 1; i <= n; i++ {
		v, err := r.Float(fmt.Sprintf("Число %d: ", i))
		if err != nil {
			return nil, err
		}
		nums = append(nums, v)
	}
	return nums, nil
}

func (r *Reader) Ints(n int) ([]int, error) {
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		v, err := r.Int(fmt.Sprintf("Число %d: ", i))
		if err != nil {
			return nil, err
		}
		nums = append(nums, v)
	}
	return nums, nil
}

var ErrNoInput = errors.New("ввод закончился")

func wrap(msg string, err error) error {
	if errors.Is(err, io.EOF) {
		return ErrNoInput
	}
	return fmt.Errorf("%s: %w", msg, err)
}

func FormatNumber(v float64) string {
	rounded := math.Round(v*1e9) / 1e9
	if rounded == 0 {
		return "0"
	}
	return strconv.FormatFloat(rounded, 'f', -1, 64)
}

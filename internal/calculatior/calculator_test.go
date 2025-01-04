package calculator

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		error      bool
	}{
		// Простые случаи
		{"1+1", 2, false},
		{"2-1", 1, false},
		{"2*3", 6, false},
		{"6/2", 3, false},
		//  Пробелы
		{" 2 + 3 * 4 ", 14, false},
		// Скобки
		{"(2+3)*4", 20, false},
		{"2+(3*4)", 14, false},
		{"(2+3)*(4+5)", 45, false},
		// Ошибочные выражения
		{"(2+3", 0, true},
		{"2+3)", 0, true},
		{"2/0", 0, true},
		{"2++3", 0, true},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		result, err := Calc(tt.expression)
		if tt.error {
			if err == nil {
				t.Errorf("Calc(%q) expected an error but got none", tt.expression)
			}
		} else {
			if err != nil {
				t.Errorf("Calc(%q) returned an unexpected error: %v", tt.expression, err)
			} else if result != tt.expected {
				t.Errorf("Calc(%q) = %v, expected %v", tt.expression, result, tt.expected)
			}
		}
	}
}

func TestMiniCalc(t *testing.T) {
	tests := []struct {
		nums      []float64
		operators []rune
		expected  []float64
		error     bool
	}{
		// Успешные вычисления
		{[]float64{2, 3}, []rune{'+'}, []float64{5}, false},
		{[]float64{5, 3}, []rune{'-'}, []float64{2}, false},
		{[]float64{2, 3}, []rune{'*'}, []float64{6}, false},
		{[]float64{6, 2}, []rune{'/'}, []float64{3}, false},
		// Ошибки
		{[]float64{6, 0}, []rune{'/'}, nil, true},
		{[]float64{2}, []rune{'+'}, nil, true},
		{[]float64{}, []rune{'+'}, nil, true},
		{[]float64{2, 3}, []rune{}, nil, true},
	}

	for _, tt := range tests {
		nums := tt.nums
		operators := tt.operators
		err := MiniCalc(&nums, &operators)

		if tt.error {
			if err == nil {
				t.Errorf("MiniCalc(%v, %v) expected an error but got none", tt.nums, tt.operators)
			}
		} else {
			if err != nil {
				t.Errorf("MiniCalc(%v, %v) returned an unexpected error: %v", tt.nums, tt.operators, err)
			} else if len(nums) != len(tt.expected) || nums[0] != tt.expected[0] {
				t.Errorf("MiniCalc(%v, %v) = %v, expected %v", tt.nums, tt.operators, nums, tt.expected)
			}
		}
	}
}

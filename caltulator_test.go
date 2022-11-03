package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 2, 3},
		{14, 2, 3},
		{0, 6, 2},
		{9, 0, 2},
		{-4, 2, -2},
	}

	for _, table := range tables {
		if _, err := Add(table.x, table.y); err != nil {
			t.Fatal("could not calculate")
		}
	}
}

func TestSubtract(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 2, -1},
		{2, 1, 1.0},
		{3, 7, -1},
		{0, 4, -2},
		{1, 0, 5},
		{-4, 5, -6},
	}

	for _, table := range tables {
		if _, err := Subtract(table.x, table.y); err != nil {
			t.Fatal("could not calculate")
		}
	}
}
func TestMultiply(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 2, 9},
		{1, 3, 2},
		{0, 6, 0},
		{4, 0, 0},
		{-4, 5, -8},
	}

	for _, table := range tables {
		if _, err := Multiply(table.x, table.y); err != nil {
			t.Fatal("could not calculate")
		}
	}
}

func TestDivide(t *testing.T) {
	tables := []struct {
		x int
		y int
		n float64
		e string
	}{
		{1, 2, 0.5, ""},
		{1.0, 2.0, 0.5, ""},
		{0, 2.0, 0, ""},
		{-4, 2.0, -2.0, ""},
	}

	for _, table := range tables {
		if total, err := Divide(table.x, table.y); err != nil && err.Error() != table.e {
			t.Errorf("Divide of (%d/%d) produced wrong error, got %v, want %v.", table.x, table.y, err.Error(), table.e)
		} else if total != table.n {
			t.Errorf("Divide of (%d/%d) produced wrong result, got %d, want %f.", table.x, table.y, total, table.n)
		}
	}
}

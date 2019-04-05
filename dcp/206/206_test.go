package dcp206

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	input := []*item{
		&item{
			Value: "a",
		},
		&item{
			Value: "b",
		},
		&item{
			Value: "c",
		},
	}
	expected := []*item{
		&item{
			Value: "c",
		},
		&item{
			Value: "b",
		},
		&item{
			Value: "a",
		},
	}
	perm := []int{2, 1, 0}

	verify(input, expected, perm, t)
}

func TestSolution2(t *testing.T) {
	input := []*item{
		&item{
			Value: "a",
		},
		&item{
			Value: "b",
		},
		&item{
			Value: "c",
		},
		&item{
			Value: "d",
		},
	}
	expected := []*item{
		&item{
			Value: "c",
		},
		&item{
			Value: "d",
		},
		&item{
			Value: "a",
		},
		&item{
			Value: "b",
		},
	}
	perm := []int{2, 3, 0, 1}

	verify(input, expected, perm, t)
}

func verify(input []*item, expected []*item, perm []int, t *testing.T) {
	result, err := solution(input, perm)
	if err != nil {
		t.Errorf("206, for input [%v], perm [%v] -> error  %v", input, perm, err)
	}
	for index, r := range result {
		e := expected[index]
		if r.Value != expected[index].Value {
			t.Errorf("206, for index [%d] = %v, want %v", index, r.Value, e.Value)
		}
	}
}

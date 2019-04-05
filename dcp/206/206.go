package dcp206

import (
	"fmt"
)

type item struct {
	Value interface{}
}

func solution(input []*item, perm []int) ([]*item, error) {
	if len(input) != len(perm) {
		return nil, fmt.Errorf("input size (%d) != perm size (%d)", len(input), len(perm))
	}
	return input, nil
}

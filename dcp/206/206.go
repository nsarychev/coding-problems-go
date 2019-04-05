package dcp206

import (
	"fmt"
)

type item struct {
	Value interface{}
}

func solution(input []*item, perm []int) ([]*item, error) {
	result := make([]*item, len(input))
	if len(input) != len(perm) {
		return nil, fmt.Errorf("input size (%d) != perm size (%d)", len(input), len(perm))
	}
	for index, p := range perm {
		result[index] = input[p]
		result[p] = input[index]
	}
	return result, nil
}

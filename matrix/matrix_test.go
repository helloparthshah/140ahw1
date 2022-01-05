package matrix

import (
	"reflect"
	"testing"
)

func TestAreAdjacent(t *testing.T) {
	tests := []struct {
		lst      []int
		a, b     int
		expected bool
	}{
		{nil, 1, 1, false},
		{[]int{}, 1, 1, false},
		{[]int{1}, 1, 1, false},
		{[]int{1, 2, 3}, 1, 3, false},
		{[]int{1, 2, 3}, 3, 1, false},
		{[]int{1, 2, 3}, 1, 2, true},
		{[]int{1, 2, 3}, 2, 3, true},
		{[]int{1, 2, 3}, 3, 2, true},
		{[]int{1, 2, 3}, 2, 1, true},
		{[]int{1, 2, 1}, 1, 1, false},
		{[]int{1, 2, 1}, 1, 2, true},
		{[]int{1, 2, 1}, 1, 4, false},
		{[]int{1, 2, 1, 4}, 1, 4, true},
		// TODO add more tests for 100% test coverage
	}
	for i, test := range tests {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("#%d: AreAdjacent(%v, %v, %v) panics: %s", i,
						test.lst, test.a, test.b, err)
				}
			}()
			actual := AreAdjacent(test.lst, test.a, test.b)
			if actual != test.expected {
				t.Errorf("#%d: AreAdjacent(%v, %v, %v)=%t; want %t", i,
					test.lst, test.a, test.b, actual, test.expected)
			}
		}()
	}
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		matrix, expected [][]int
	}{
		{nil, nil},
		{[][]int{}, [][]int{}},
		{[][]int{{1, 2, 3, 4}}, [][]int{{1}, {2}, {3}, {4}}},
		{[][]int{{1}, {2}, {3}, {4}}, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}}},
		{[][]int{{1, 3}, {2, 4}}, [][]int{{1, 2}, {3, 4}}},
		// TODO add more tests for 100% test coverage
	}
	for i, test := range tests {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("#%d: Transpose(%v) panics: %s", i,
						test.matrix, err)
				}
			}()
			actual := Transpose(test.matrix)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("#%d: Transpose(%v)=%v; want %v", i,
					test.matrix, actual, test.expected)
			}
		}()
	}
}

func TestAreNeighbors(t *testing.T) {
	type Test struct {
		matrix   [][]int
		a, b     int
		expected bool
	}
	tests := []Test{
		{nil, 1, 2, false},
		{[][]int{}, 1, 2, false},
		{[][]int{{1, 2, 3}}, 1, 2, true},
		{[][]int{{1, 2, 3}}, 1, 3, false},
		{[][]int{{1}, {2}, {3}}, 1, 2, true},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 1, 2, true},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 2, 6, false},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 1, 6, false},
		// TODO if needed, add more tests for 100% test coverage
	}
	for i, test := range tests {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("#%d: AreNeighbors(%v, %v, %v) panics: %s", i,
						test.matrix, test.a, test.b, err)
				}
			}()
			actual := AreNeighbors(test.matrix, test.a, test.b)
			if actual != test.expected {
				t.Errorf("#%d: AreNeighbors(%v, %v, %v)=%t; want %t", i,
					test.matrix, test.a, test.b, actual, test.expected)
			}
		}()
	}
}

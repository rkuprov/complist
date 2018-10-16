package complist

import (
	"fmt"
	"testing"
)

func TestIsSame(t *testing.T) {
	table := []struct {
		a   interface{}
		b   interface{}
		ans bool
	}{
		{"home", "ohme", true},
		{"home", "emoh", true},
		{"balaclava", "bllaacava", true},
		{"doggo", "odgog", true},
		{"anagram", "analytical", false},
		{"sitting", "stingin", false},
		{"string", "stirng", true},
		{"ocean", "Ocean", false},
		{"122345", "12345", false},
		{"122345", "1223345", false},
		{"12345", "54321", true},
		{int16(1234), int16(4321), true},
		{122345, 1223345, false},
		{12.2345, 1223.345, false},
		{12.23345, 12233.45, false},
		{float32(12.23), float32(122.3), false},
		{int64(122345), int64(12345), false},
		{int64(122345), int64(1235), false},
		{uint16(2345), uint16(5432), true},
		{int32(12345111), int32(54321111), true},
		{uint32(12345111), uint32(54321111), true},
		{uint64(12345111), uint64(54321111), true},
		{[]string{"a", "b", "c"}, []string{"a", "bc"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]int{1, 1, 2}, []int{1, 2, 3}, false},
		{[]int{1, 1, 2}, []int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1}, []int{1, 2, 3}, false},
		{[]float64{1.2, 2.37, 3}, []float64{1.2, 2.37, 3}, true},
		{map[int]int{1: 1, 2: 3, 3: 25}, map[int]int{1: 1, 2: 3, 3: 25}, true},
		{map[string]int{"1": 1, "2": 3, "3": 25}, map[string]int{"1": 1, "2": 3, "3": 25}, true},
		{map[string]int{"1": 1, "2": 3, "3": 25}, map[string]int{"1": 1, "3": 25, "2": 3}, true},
		{map[int]int{1: 1, 2: 3, 3: 25}, map[int]int{1: 1, 2: 3, 3: 25, 4: 13}, false},
		{map[string]int{"1": 1, "2": 3, "3": 25}, map[string]int{"1": 1, "2": 3, "3": 25, "4": 13}, false},
		{[3]int{1, 1, 2}, [3]int{1, 2, 3}, false},
		{[3]int{1, 1, 2}, [4]int{1, 2, 3, 4}, false},
		{[3]int{1, 1, 1}, [3]int{1, 2, 3}, false},
		{[3]int{1, 2, 3}, [3]int{1, 2, 3}, true},
		{[3]int{1, 2, 3}, [3]int{1, 3, 2}, true},
		{[3]string{"1", "string", "$#"}, [3]string{"string", "$#", "1"}, true},
		{[3]bool{true, true, false}, [3]bool{false, true, true}, true},
	}

	for _, line := range table {
		issame, err := IsSame(line.a, line.b)
		if err != nil {
			fmt.Println("error encountered: ", err)
		}
		if issame != line.ans {
			t.Errorf("IsSame error for : %v & %v. Want %v, got %v", line.a, line.b, line.ans, issame)
		}
	}
}

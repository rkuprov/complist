package complist

import (
	"fmt"
	"reflect"
	"strconv"
)

//IsSame checks whether two items have an identical collection of elements
func IsSame(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false, fmt.Errorf("element type mismatch: %v doesn't match %v", reflect.TypeOf(a), reflect.TypeOf(b))
	}
	//Values to pass through the switch statement
	vala := reflect.ValueOf(a)
	valb := reflect.ValueOf(b)
	//only one item.Kind is needed since both need to be identical to reach this point
	switch vala.Kind() {
	case reflect.String:
		return IsSameString(vala.String(), valb.String()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return IsSameNumber(getIntDigits(int(vala.Uint())), getIntDigits(int(valb.Uint()))), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return IsSameNumber(getIntDigits(int(vala.Int())), getIntDigits(int(valb.Int()))), nil
	case reflect.Float32, reflect.Float64:
		return false, fmt.Errorf("can't analyze floats")
	case reflect.Array:
		return IsSameArray(vala, valb), nil
	case reflect.Map:
		return IsSameMap(vala, valb), nil
	case reflect.Slice:
		return IsSameSlice(vala, valb), nil
	case reflect.Struct:
		return false, fmt.Errorf("can't analyze structs... yet")
	}
	return false, nil
}

//IsSameString compares strings
func IsSameString(a, b string) bool {
	al := len(a)
	bl := len(b)
	if al != bl {
		return false
	}
	//setting up memory maps for comparison
	ma := make(map[byte]int)
	mb := make(map[byte]int)
	//populating both maps
	for i := 0; i < al; i++ {
		ma[a[i]]++
	}
	for i := 0; i < bl; i++ {
		mb[b[i]]++
	}
	//comparing one map to the other. This is ok, since they're the same length.
	for k, v := range ma {
		if mb[k] != v {
			return false
		}
	}
	return true
}

//IsSameNumber checks whether two numbers have the same collection of digits.
func IsSameNumber(a, b []int) bool {
	al := len(a)
	bl := len(b)
	if al != bl {
		return false
	}
	//setting up memory maps for comparison
	ma := make(map[int]int)
	mb := make(map[int]int)
	//populating the maps
	for i := 0; i < al; i++ {
		ma[a[i]]++
	}
	for i := 0; i < bl; i++ {
		mb[b[i]]++
	}
	//comparing the two maps against each other.
	for k, v := range ma {
		if mb[k] != v {
			return false
		}
	}

	return true
}

//IsSameSlice compares two slices according to the collection of their elements
func IsSameSlice(a, b reflect.Value) bool {
	al := a.Len()
	bl := b.Len()
	if al != bl {
		return false
	}
	//setting up memory maps for comparison
	ma := make(map[interface{}]int)
	mb := make(map[interface{}]int)
	//populating the maps
	for i := 0; i < al; i++ {
		ma[a.Index(i).Interface()]++
	}
	for i := 0; i < bl; i++ {
		mb[b.Index(i).Interface()]++
	}
	//comparing the two maps against each other.
	for k, v := range ma {
		if mb[k] != v {
			return false
		}
	}
	return true
}

//IsSameMap compares two maps element by element.
func IsSameMap(a, b reflect.Value) bool {
	//making sure both maps have the same number of elements.
	if a.Len() != b.Len() {
		return false
	}
	//only one map is needed for checking. Any discrepancies will fail the test.
	ak := a.MapKeys() //getting the keys for map 'a'.
	//Comparing the two maps with the key from map 'a'.
	for _, key := range ak {
		if a.MapIndex(key).Interface() != b.MapIndex(key).Interface() {
			return false
		}
	}
	return true
}

//IsSameArray compares two arrays according to the collection of their elements
func IsSameArray(a, b reflect.Value) bool {
	//setting up memory maps for comparison
	ma := make(map[interface{}]int)
	mb := make(map[interface{}]int)
	//populating the maps
	for i := 0; i < a.Len(); i++ {
		ma[a.Index(i).Interface()]++
	}
	for i := 0; i < b.Len(); i++ {
		mb[b.Index(i).Interface()]++
	}
	//comparing the two maps against each other.
	for k, v := range ma {
		if mb[k] != v {
			return false
		}
	}
	return true
}

func getIntDigits(n int) []int {
	var ret []int
	//could've converted to string and then interpreted each digit, but this was more fun.
	if n < 10 {
		return append(ret, n)
	}
	ret = append(ret, n%10)
	n = n / 10

	return append(ret, getIntDigits(n)...)
}

func getUIntDigits(n uint) []uint {
	var ret []uint
	//Same as with int.
	if n < 10 {
		return append(ret, n)
	}
	ret = append(ret, n%10)
	n = n / 10

	return append(ret, getUIntDigits(n)...)

}

//non-function code at this point. Floats act funny when you shift decimals around.
func getFloat64Digits(n float64) []int {
	s := fmt.Sprint(n)
	var ret []rune
	for _, char := range s {
		if char >= '0' && char <= '9' {
			ret = append(ret, char)
		}
	}
	ss, _ := strconv.Atoi(string(ret))
	return getIntDigits(ss)
}

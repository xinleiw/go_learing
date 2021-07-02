package deep_equal

import (
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	// t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := [...]int{1, 2, 3, 4, 5}

	// t.Log(l1 == l2)
	t.Log(reflect.DeepEqual(l1, l2))
	t.Log(reflect.DeepEqual(l1, l3))

	s1 := "aaaa"
	s2 := "aaaa"
	t.Log(s1 == s2)
}

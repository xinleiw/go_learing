package map_ext

import "testing"


func TestMapWithFunValue(t *testing.T){
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op * op}
	m[3] = func(op int) int {return op * op * op}
	t.Log(m[1](2), m[2](2), m[3](2))
	t.Log(m[1](4))
}

func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
		mySet[1] = true
		n := 3
		if mySet[n]{
			t.Logf("%d is existing", n)
		}else {
			mySet[n] = true
			t.Logf("%d is not existing", n)
		}
		t.Log(len(mySet))
		delete(mySet, 1)
		if mySet[1]{
			t.Logf("%d is existing", 1)
		}else {
			t.Logf("%d is not existing", 1)
		}
		for k,v := range mySet{
			t.Log(k, v)
		}
}
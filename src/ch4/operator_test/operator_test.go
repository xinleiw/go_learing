package operator_test
import "testing"

const(
	Readable = 1 << iota
	Writable
	Executeable
)

func TestCompareArrary(t * testing.T){
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d) 
	// t.Log(b > a)
}

func TestBitClear(t *testing.T){
	a:=7
	a = a&^Readable
	t.Log(a)
	t.Log(a&Readable, a&Writable, a&Executeable)
	t.Log(Readable, Writable, Executeable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executeable == Executeable)
}
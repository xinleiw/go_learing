package constant_test
import "testing"

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1 << iota
	Writable
	Executeable
)

func TestConstantTry(t *testing.T){
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T){
	a:=1
	t.Log(a&Readable, a&Writable, a&Executeable)
	t.Log(Readable, Writable, Executeable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executeable == Executeable)
}
package type_test

import "testing"

type MyInt int32

func TestImpliccitConvert(t *testing.T){
	var a int32 = 1
	var b int64
	var c MyInt
	// b = a
	b = int64(a)
	c = MyInt(a)
	// c = a
	t.Log(b, a, c)
}

func TestPoint(t *testing.T){
	a:=1
	aPtr := &a
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}


func TestSring(t *testing.T){
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == ""{
		t.Log("字符串初始化完成")
	}
}

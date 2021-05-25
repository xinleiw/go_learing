package string_fn

import ("testing"
		"strings"
	    "strconv")

func TestStringFn(t *testing.T){
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	for _, item := range parts{
		t.Log(item)
	}
	t.Log(strings.Join(parts, "_"))
}

func TestConv(t *testing.T){
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// t.Log(10 + strconv.Atoi("1"))
	if i, err := strconv.Atoi("10"); err==nil{
		t.Log(10 + i)
	}
}
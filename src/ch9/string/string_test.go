package string_test
import "testing"

func TestString(t *testing.T){
	var s string
	t.Log(s)
	s = "hello"
	t.Log(s)
	t.Log(len(s))
	// string是不可变的
	// s[1] = "3"
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "我是中国人"
	t.Log(len(s))
	s = "中"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UFT8 %x", s)
}

func TestStringToRune(t *testing.T){
	s := "中华人民共和国"
	for _, item := range s{
		t.Logf("%c", item)
	}
}
package interface_test

import "testing"

type Programmer interface {
	WriteHelloworld() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloworld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloworld())
}

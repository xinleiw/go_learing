package loop_test

import "testing"

func TestWhileLoop(t *testing.T){
	 
	for n := 0; n < 5; n++ {
		t.Log(n)
		
	}
}


func TestWhileLoop1(t *testing.T){
	n := 0 
	for n < 5{
		t.Log(n)
		n++
	}
}
package condition_test

import ("testing"
		"fmt")

func TestIfMultiSec(t *testing.T){
	if a:=1 == 1; a{
		t.Log("1 == 1")
	}
}

func TestSwitchMultiSec(t *testing.T){
	for i:=0;i<5;i++{
		fmt.Printf("%d", i)
		switch i {
		case 0, 2:
			t.Logf("%v is Even", i)
		case 1, 3:
			t.Logf("%v is Odd", i)
		default:
			t.Logf("%v is not 0-3", i)
			
		}
	}
}

func TestSwitchCaseCondition(t *testing.T){
	for i:=0;i<5;i++{
		fmt.Printf("%d", i)
		switch {
		case i % 2 == 0:
			t.Logf("%v is Even", i)
		case i % 2 == 1:
			t.Logf("%v is Odd", i)
		default:
			t.Logf("%v unkonw", i)
			
		}
	}
}
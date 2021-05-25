package map_test

import "testing"

func TestInitMap(t *testing.T){
	m1 := map[int]int{1:1,2:4,3:6}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 8
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T){
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// key 不存在的时候 ok 为false 存在则为true
	// m1[3] = 0
	if v, ok :=m1[3];ok{
		t.Logf("key 3's value is %d", v)
	}else{
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T){
	m1 := map[int]int{1:1,2:4,3:6}
	for k, v := range m1{
		t.Log(k, v)
	}
}
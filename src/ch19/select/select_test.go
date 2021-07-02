package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

// func TestService(t *testing.T) {
// 	fmt.Println(service())
// 	otherTask()
// }

func AsyncService() chan string {
	// var wg sync.WaitGroup
	retCh := make(chan string, 1)
	// wg.Add(1)
	go func() {
		ret := service()
		fmt.Println("Return result.")
		retCh <- ret
		fmt.Println("service exied.")
		// wg.Done()
	}()
	// wg.Wait()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	default:
		fmt.Println("default")
	}

}

package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroitine(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// go func() {
		// 	fmt.Println(i)
		// }()
	}
	time.Sleep(time.Second * 5)
}

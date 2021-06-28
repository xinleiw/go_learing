package panic_recover_test

import (
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	fmt.Println("Start")
	// panic(errors.New("Something error"))
	os.Exit(-1)

}

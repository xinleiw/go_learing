package err_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargeThanHunderedError = errors.New("n should be not large than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargeThanHunderedError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil

}

func TestGetFibonacci(t *testing.T) {
	// var res []int = GetFibonacci(10)
	if res, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Printf("%v", LessThanTwoError)
		}
		t.Error(err)
	} else {
		t.Log(res)
	}
	res, err := GetFibonacci(-10)
	t.Log(res, err)
	// fmt.Printf("res:%v", res)
	GetFibonacci(-10)
}

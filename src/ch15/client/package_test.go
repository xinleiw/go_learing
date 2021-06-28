package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	ret := series.GetFibonacciSeries(5)
	ret2 := series.Square(5)
	t.Log(ret)
	t.Log(ret2)
}

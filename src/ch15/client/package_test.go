package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {

	value := series.GetFibonacciSerie(10)

	t.Log(value)
}

package error

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {

	if v, err := getValue(11); err == nil {
		t.Log(v)
	} else {
		t.Log(err)
	}
}

var errorLarger = errors.New("value is larger ")
var errorless = errors.New("value is lsss")

func getValue(op int) (int, error) {
	if op < 10 {
		return 0, errorless
	}

	if op > 100 {
		return 0, errorLarger
	}

	return op, nil
}

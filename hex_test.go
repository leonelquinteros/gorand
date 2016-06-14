package gorand

import (
	"testing"
)

func TestGetHex(t *testing.T) {
	str, err := GetHex(8)
	if err != nil {
		t.Error(err.Error())
	}

	if len(str) != 16 {
		t.Error("Length of string isn't 16")
	}
}

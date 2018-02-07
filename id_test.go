package gorand

import (
	"testing"
)

func TestID(t *testing.T) {
	id, err := ID()
	if err != nil {
		t.Error(err.Error())
	}

	if len(id) != 128 {
		t.Error("Length of UUID isn't 128")
	}
}

func BenchmarkID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ID()
	}
}

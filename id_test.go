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

func TestUUID(t *testing.T) {
	uuid, err := UUID()
	if err != nil {
		t.Error(err.Error())
	}

	if len(uuid) != 36 {
		t.Error("Length of UUID isn't 36")
	}
}

func TestMarshalUUID(t *testing.T) {
	uuid, err := UUID()
	if err != nil {
		t.Error(err.Error())
	}

	m, err := UnmarshalUUID(uuid)
	if err != nil {
		t.Error(err.Error())
	}

	u, err := MarshalUUID(m)
	if err != nil {
		t.Error(err.Error())
	}

	if u != uuid {
		t.Errorf("%s != %s after Unmarshal and Marshal", u, uuid)
	}
}

func BenchmarkID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ID()
	}
}

func BenchmarkUUID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UUID()
	}
}

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

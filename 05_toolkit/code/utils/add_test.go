package utils

import (
	"testing"
)

func TestAverage(t *testing.T) {
	expected := 5
	actual := Add(1, 2, 4)

	if actual != expected {
		t.Error("Avarage was incorrect, got: ", actual, " expected: ", expected)
	}
}

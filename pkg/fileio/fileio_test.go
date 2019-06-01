package fileio

import (
	"testing"
)

func TestAverage(t *testing.T) {
	hash, err := HashFile("testdata/text.txt")
	if err != nil {
		t.Error("Error hashing file")
	}
	if hash != "b6938a42d7e7c495c12a66fa98d0b5c0919a9d8c5dbda25f6b3688268cafd3da" {
		t.Error("Hash did not equal expected value")
	}
}

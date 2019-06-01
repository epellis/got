package fileio

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
)

// Opens a file and returns the SHA 256 hash of it.
func HashFile(filepath string) (string, error) {
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(dat)

	checksum := hex.EncodeToString(hasher.Sum(nil))

	return checksum, err
}

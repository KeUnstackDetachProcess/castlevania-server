package random

import (
	"crypto/rand"
)

func GenerateByteArray(len int) ([]byte, error) {

	ret := make([]byte, len)
	_, err := rand.Read(ret)

	return ret, err
}

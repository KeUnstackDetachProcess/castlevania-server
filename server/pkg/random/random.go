package random

import (
	// "bytes"
	"crypto/rand"
	"cv-server/pkg/log"
	"encoding/hex"
)

// Generates a byte array which length depends on the "len" parameter
func GenerateByteArray(len int) []byte {
	ret := make([]byte, len) // Our return value

	_, err := rand.Read(ret) // https://pkg.go.dev/crypto/rand#Read

	if err != nil {
		log.Error("Byte array generation failed: " + err.Error())
		return nil
	}

	// Byte array generation was successful
	log.Success("Byte array generation OK: " + hex.EncodeToString(ret[:]))

	return ret
}




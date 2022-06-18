package api

import (
	"crypto/sha256"
	"cv-client/pkg/log"
	"encoding/hex"
	"io"
	"os"
)

type User struct {
}

type PatchShield struct {
}

const (
	CONNECT_REFUSE = ""
	CONNECT_ERROR  = ""
)

func Connect(endpoint string) error {

	return nil
}

func (state *PatchShield) MakeSafetyReport() {

	// Open current executable file
	f, err := os.Open(os.Args[0])
	if err != nil {
		log.Error(err.Error())
	}

	defer f.Close()

	// Get current executable file SHA256 hash checksum
	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Error(err.Error())
	}

	value := hex.EncodeToString(hasher.Sum(nil))

	// Validate checksum
	if value != CLIENT_CHECKSUM {
		log.Warning("Client checksum is different, the software could be compromised! Be careful!")
	}

	// Get server checksum and validate it

	return true
}

func (user User) CreateNewUser() User {

	return user
}

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

type LocalReport struct {
	onion      bool
	corruption bool
}

type RemoteReport struct {
	weak       bool
	encryption bool
	corruption bool
}

type SecurityReport struct {
	local  LocalReport
	remote RemoteReport
}

func Connect(endpoint string) error {

	return nil
}

func (sr *SecurityReport) MakeSecurityReport() {

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
		sr.local.corruption = true
	}

	// -- temporarely hard-coded --

	// Ensure client onion routing is enabled / works properly
	sr.local.onion = true

	// Get server checksum and validate it
	sr.remote.corruption = false
	// Get server encryption (uses E2E) state
	sr.remote.encryption = false
	// Get server weakness state (compares default db psw hash)
	sr.remote.weak = false
}

func (user User) CreateNewUser() User {

	return user
}

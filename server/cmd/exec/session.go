package main

import (
	"cv-server/pkg/log"
	"cv-server/pkg/rsa"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type SessionCertificate struct {
	Ciphertext []byte `json:"c"`
	Signature  []byte `json:"s"`
	PublicKey  []byte `json:"p"`
}

type Response struct {
	Ok    bool   `json:"ok"`
	Token []byte `json:"t"`
}

func validateSession(rw http.ResponseWriter, r *http.Request) {

	var ok bool = true
	var token []byte = []byte{}
	var sessionCert SessionCertificate

	if body, err := io.ReadAll(r.Body); err != nil {
		ok = false
	} else if err = json.Unmarshal(body, &sessionCert); err != nil {
		ok = false
	} else if pubKey, err := rsa.DecodePublicKey(sessionCert.PublicKey); err != nil {
		ok = false
	} else if rsa.Validate(sessionCert.Ciphertext, sessionCert.Signature, pubKey); err != nil {
		ok = false
	}

	if ok {
		token = generateToken()
	}

	response, _ := json.Marshal(Response{
		Ok:    ok,
		Token: token,
	})

	log.Response(rw, string(response))
}

func InitializeSessionRoutes(router *mux.Router) {
	// validate session id signature and provide an unique token
	router.HandleFunc("/api/session/validate", validateSession).Methods("POST")
}

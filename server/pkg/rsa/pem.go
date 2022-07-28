package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Public Keys

func EncodePublicKey(pubKey *rsa.PublicKey) []byte {

	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pubKey),
		},
	)
}

func DecodePublicKey(pemKey []byte) (*rsa.PublicKey, error) {

	block, _ := pem.Decode(pemKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// Private Keys

func EncodePrivateKey(privKey *rsa.PrivateKey) []byte {

	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privKey),
		},
	)
}

func DecodePrivateKey(pemKey []byte) (*rsa.PrivateKey, error) {

	block, _ := pem.Decode(pemKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

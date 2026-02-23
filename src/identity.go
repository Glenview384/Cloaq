// NOTICE
// Project Name: Cloaq
// Copyright © 2026 Neil Talap and/or its designated Affiliates.

package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/hex"
)

type Identity struct {
	PrivateKey *ecdh.PrivateKey
	PublicKey  *ecdh.PublicKey
}

func (i *Identity) String() string {
	return hex.EncodeToString(i.PublicKey.Bytes())
}

func GenerateIdentity() (*Identity, error) {
	identity := &Identity{}
	pKey, err := ecdh.X25519().GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	identity.PrivateKey = pKey
	identity.PublicKey = pKey.Public().(*ecdh.PublicKey)

	return identity, nil
}

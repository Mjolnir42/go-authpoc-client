/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"hash"
	"math/big"
	"net"
	"time"

	"golang.org/x/crypto/nacl/box"

	"github.com/dchest/blake2b"
	"github.com/satori/go.uuid"
)

type AuthRequest struct {
	Public               string    `json:"public"`
	Request              uuid.UUID `json:"request"`
	InitializationVector string    `json:"initialization_vector"`
	Token                string    `json:"token,omitempty"`
	private              string    `json:"-"`
	peer                 string    `json:"-"`
	sourceIP             net.IP    `json:"-"`
	count                uint      `json:"-"`
	time                 time.Time `json:"-"`
}

func NewAuthRequest() *AuthRequest {
	var (
		err                   error
		hBlake                hash.Hash
		bRandom, bSecret, bIV []byte
		publicKey, privateKey *[32]byte
	)
	a := AuthRequest{}

	// generate the IV from 192bit random
	bIV = make([]byte, 24)
	if _, err = rand.Read(bIV); err != nil {
		return nil
	}
	a.InitializationVector = hex.EncodeToString(bIV)

	// generate keypair, starting with 1024bit random
	bRandom = make([]byte, 128)
	if _, err = rand.Read(bRandom); err != nil {
		return nil
	}
	// hashed down to 256 bit
	hBlake = blake2b.New256()
	hBlake.Write(bRandom)
	bSecret = hBlake.Sum(nil)
	// generate keys
	if publicKey, privateKey, err = box.GenerateKey(
		bytes.NewReader(bSecret),
	); err != nil {
		return nil
	}
	a.Public = hex.EncodeToString(publicKey[:])
	a.private = hex.EncodeToString(privateKey[:])
	return &a
}

func (a *AuthRequest) IsExpired() bool {
	return time.Now().UTC().After(a.time.UTC().Add(60 * time.Second))
}

func (a *AuthRequest) SameSource(ip net.IP) bool {
	return a.sourceIP.Equal(ip)
}

// Nonces are built by interpreting the IV as a positive integer
// number and adding the count of requested nonces; thus implementing
// a simple counter. The IV itself is never used as a nonce. Returns
// nil on error.
func (a *AuthRequest) NextNonce() *[24]byte {
	var (
		ib []byte
		e  error
	)

	a.count += 1
	if ib, e = hex.DecodeString(a.InitializationVector); e != nil {
		return nil
	}
	iv := big.NewInt(0)
	iv.SetBytes(ib)
	iv.Abs(iv)
	iv.Add(iv, big.NewInt(int64(a.count)))
	if len(iv.Bytes()) != 24 {
		return nil
	}

	nonce := &[24]byte{}
	copy(nonce[:], iv.Bytes()[0:24])
	return nonce
}

func (a *AuthRequest) PeerKey() *[32]byte {
	var (
		pk []byte
		e  error
	)
	if pk, e = hex.DecodeString(a.peer); e != nil {
		return nil
	}
	if len(pk) != 32 {
		return nil
	}
	peer := &[32]byte{}
	copy(peer[:], pk[0:32])
	return peer
}

func (a *AuthRequest) SetPeerKey(k *[32]byte) {
	a.peer = hex.EncodeToString(k[:])
}

func (a *AuthRequest) PrivateKey() *[32]byte {
	var (
		pk []byte
		e  error
	)
	if pk, e = hex.DecodeString(a.private); e != nil {
		return nil
	}
	if len(pk) != 32 {
		return nil
	}
	private := &[32]byte{}
	copy(private[:], pk[0:32])
	return private
}

func (a *AuthRequest) PublicKey() *[32]byte {
	var (
		pk []byte
		e  error
	)
	if pk, e = hex.DecodeString(a.Public); e != nil {
		return nil
	}
	if len(pk) != 32 {
		return nil
	}
	public := &[32]byte{}
	copy(public[:], pk[0:32])
	return public
}

func (a *AuthRequest) SetRequest(s string) error {
	var err error

	if a.Request, err = uuid.FromString(s); err != nil {
		return err
	}

	return nil
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

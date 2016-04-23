/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import "golang.org/x/crypto/nacl/box"

func decryptBytes(a *AuthRequest, cipherText []byte) *[]byte {

	var (
		nonce            *[24]byte
		peerKey, privKey *[32]byte
		plainText        []byte
		ok               bool
	)

	if nonce = a.NextNonce(); nonce == nil {
		return nil
	}

	if peerKey = a.PeerKey(); peerKey == nil {
		return nil
	}

	if privKey = a.PrivateKey(); privKey == nil {
		return nil
	}

	if plainText, ok = box.Open(nil, cipherText, nonce, peerKey, privKey); !ok {
		return nil
	}
	return &plainText
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

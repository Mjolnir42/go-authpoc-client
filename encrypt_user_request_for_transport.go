/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func encryptUserRequestForTransport(u UserRequest, a *AuthRequest) ([]byte, error) {
	var (
		plainText  []byte
		cipherText *[]byte
		transport  string
		err        error
	)

	if plainText, err = json.Marshal(u); err != nil {
		return nil, err
	}

	if cipherText = encryptBytes(a, plainText); cipherText == nil {
		return nil, fmt.Errorf("Encryption of user request failed")
	}

	transport = base64.StdEncoding.EncodeToString(*cipherText)

	return []byte(transport), nil
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

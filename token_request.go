/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import "fmt"

func tokenRequest(cipherText []byte, auth *AuthRequest) (UserRequest, error) {
	var (
		url string
	)

	url = fmt.Sprintf(
		"%s/%s/%s",
		ApiAddress,
		"authenticate/token",
		auth.Request,
	)

	return serverRequest(url, cipherText, auth)
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

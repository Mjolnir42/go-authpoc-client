/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func cmdActionAuthenticate(c *cli.Context) {
	var (
		user, password string
		err            error
		auth           *AuthRequest
		uReq           UserRequest
		cipherText     []byte
	)

	// get username
	if c.IsSet("user") {
		user = c.String("user")
	} else {
		user = readConfirmedUsername()
	}

	// get password
	// aborts on error
	password = readInputPassword()

	// initiate auth request
	auth = NewAuthRequest()

	if err = initiateAuthentication(auth); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	// encrypt request with negotiated key
	uReq = UserRequest{
		User:     user,
		Password: password,
	}
	if cipherText, err = encryptUserRequestForTransport(uReq, auth); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	// request token
	if uReq, err = tokenRequest(cipherText, auth); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	// validate received password token
	if err = validatePasswordToken(user, uReq.Token); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Token: %s\n", uReq.Token)
	fmt.Printf("Expires at: %s\n", uReq.ExpiresAt)
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

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

func cmdActionRegister(c *cli.Context) {
	var (
		err            error
		user, password string
		happy          bool
		cipherText     []byte
	)

	/*
	 * GET CREDENTIALS
	 */
	// read in username
	user = readConfirmedUsername()

password_read:
	// read in password from terminal
	password = readVerifiedPassword()

	// evaluate password
	if happy, err = evaluatePassword(3, password, user); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !happy {
		password = ""
		goto password_read
	}

	/*
	 * SEND AUTH INITIATOR REQUEST TO SERVER
	 */
	auth := NewAuthRequest()

	if err := initiateAuthentication(auth); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	/*
	 * ENCRYPT USER REQUEST WITH COMMON KEY
	 */
	// construct the request
	uReq := UserRequest{
		User:     user,
		Password: password,
	}
	// encrypt it
	if cipherText, err = encryptUserRequestForTransport(uReq, auth); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	/*
	 * SEND REGISTER REQUEST TO SERVER
	 */
	if uReq, err = registerRequest(cipherText, auth); err != nil {
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

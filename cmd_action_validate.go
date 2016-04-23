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

func cmdActionValidate(c *cli.Context) {
	var (
		user, token string
		err         error
	)

	// get username
	if c.IsSet("user") {
		user = c.String("user")
	} else {
		user = readConfirmedUsername()
	}

	// get password token
	if c.IsSet("token") {
		token = c.String("token")
	} else {
		token = readConfirmedToken()
	}

	// validate
	if err = validatePasswordToken(user, token); err != nil {
		fmt.Fprintf(
			os.Stderr,
			RED+FAILURE+CLEAR+" %s %s",
			"Verification failed:",
			err.Error(),
		)
		os.Exit(1)
	}
	fmt.Fprintf(
		os.Stderr,
		GREEN+SUCCESS+CLEAR+" %s\n",
		"Token successfully verified",
	)
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

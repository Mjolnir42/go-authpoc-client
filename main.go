/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const (
	SUCCESS = "\xe2\x9c\x94"
	FAILURE = "\xe2\x9c\x98"

	RED   = "\x1b[31m"
	GREEN = "\x1b[32m"
	CLEAR = "\x1b[0m"

	ApiAddress = "https://localhost:9999"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-auth-example-client"
	app.Usage = "Example/POC GoLang cli authentication"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:   "register",
			Usage:  "Register a new account",
			Action: cmdActionRegister,
		},
		{
			Name:   "authenticate",
			Usage:  "Authenticate against an account",
			Action: cmdActionAuthenticate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Username to use for authentication",
				},
			},
		},
		{
			Name:   "validate",
			Usage:  "Validate an authentication token",
			Action: cmdActionValidate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Username to use for authentication",
				},
				cli.StringFlag{
					Name:  "token, t",
					Usage: "Token to use for authentication",
				},
			},
		},
	}
	app.Run(os.Args)
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

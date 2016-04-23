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
	"strings"

	"github.com/peterh/liner"
)

// readUsername reads an input string from the terminal. Returns
// error if interrupted
func readUsername() (string, error) {
	var (
		err  error
		user string
	)

	const prompt string = `Enter username: `

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	if user, err = line.Prompt(prompt); err != nil {
		return "", err
	}

	return user, nil
}

// readConfirmedUsername invokes readUsername until it has a
// non-empty input string as username. Aborts on error.
func readConfirmedUsername() string {
	var (
		user string
		err  error
	)

loop:
	for user == "" {
		if user, err = readUsername(); err != nil {
			if err == liner.ErrPromptAborted {
				os.Exit(0)
			} else {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
	// _NOTHING_ can handle the Version 6 AT&T UNIX /etc/passwd
	// field separator as part of the username. This includes
	// HTTP Basic Auth.
	if strings.Contains(user, `:`) {
		fmt.Fprintf(
			os.Stderr,
			RED+FAILURE+CLEAR+" %s\n",
			"Username must not contain the : character.",
		)
		goto loop
	}
	return user
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

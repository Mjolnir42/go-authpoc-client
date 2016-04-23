/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/peterh/liner"
)

var PasswordMismatchError = errors.New(`Passwords did not match`)

// readPassword reads a single password from the terminal. Returns
// error if the input could not be read or was interrupted.
func readPassword() (string, error) {
	var (
		err error
		p1  string
	)

	const (
		prompt1 string = `Enter password: `
	)

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	if p1, err = line.PasswordPrompt(prompt1); err != nil {
		return "", err
	}

	return p1, nil
}

// readConfirmedPassword reads a password twice from the terminal and
// checks if they match. Returns PasswordMismatchError if they do not.
func readConfirmedPassword() (string, error) {
	var (
		err    error
		p1, p2 string
	)

	const (
		prompt1 string = `Enter password: `
		prompt2 string = `Repeat password: `
	)

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	if p1, err = line.PasswordPrompt(prompt1); err != nil {
		return "", err
	}

	if p2, err = line.PasswordPrompt(prompt2); err != nil {
		return "", err
	}

	if p1 != p2 {
		return "", PasswordMismatchError
	}

	return p1, nil
}

// readVerifiedPassword runs readConfirmedPassword until a matching
// password was provided
func readVerifiedPassword() string {
	var (
		password string
		err      error
	)

password_read:
	for password == "" {
		if password, err = readConfirmedPassword(); err != nil {
			if err == liner.ErrPromptAborted {
				os.Exit(0)
			} else if err == PasswordMismatchError {
				fmt.Fprintf(
					os.Stderr,
					RED+FAILURE+CLEAR+" %s\n",
					err.Error(),
				)
				continue password_read
			} else {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
	fmt.Fprintf(
		os.Stderr,
		GREEN+SUCCESS+CLEAR+" %s\n",
		"Entered passwords match",
	)
	return password
}

// readInputPassword runs readPassword until a not empty password
// was provided
func readInputPassword() string {
	var (
		password string
		err      error
	)

	for password == "" {
		if password, err = readPassword(); err != nil {
			if err == liner.ErrPromptAborted {
				os.Exit(0)
			} else {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
	return password
}

func verifyChoice() (string, error) {
	var (
		err    error
		choice string
	)

	const (
		prompt string = `Select this password? (y/n): `
	)

	line := liner.NewLiner()

	line.SetCtrlCAborts(true)

	if choice, err = line.Prompt(prompt); err != nil {
		return "", err
	}

	return choice, nil
}

// readToken reads a single token from the terminal. Returns
// error if the input could not be read or was interrupted.
func readToken() (string, error) {
	var (
		err error
		t1  string
	)

	const (
		prompt1 string = `Enter token: `
	)

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	if t1, err = line.PasswordPrompt(prompt1); err != nil {
		return "", err
	}

	return t1, nil
}

// readConfirmedToken invokes readToken until it has a
// non-empty input string as token. Aborts on error.
func readConfirmedToken() string {
	var (
		token string
		err   error
	)

	for token == "" {
		if token, err = readToken(); err != nil {
			if err == liner.ErrPromptAborted {
				os.Exit(0)
			} else {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
	return token
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

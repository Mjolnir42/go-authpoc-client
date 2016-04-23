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

	"github.com/nbutton23/zxcvbn-go"
	"github.com/nbutton23/zxcvbn-go/scoring"
	"github.com/peterh/liner"
)

func evaluatePassword(minScore int, sl ...string) (bool, error) {
	var (
		err     error
		quality scoring.MinEntropyMatch
		choice  string
	)

	if len(sl) < 1 {
		return false, fmt.Errorf("No password given for judgement")
	}

	// second argument are additional strings whose use should be
	// punished, ie. username
	quality = zxcvbn.PasswordStrength(sl[0], sl[1:])

	// display evaluation summary report
	fmt.Printf(
		`Password score    (0-4): %d
Estimated entropy (bit): %f
Estimated time to crack: %s%s`,
		quality.Score,
		quality.Entropy,
		quality.CrackTimeDisplay, "\n",
	)

	// enforce chance for a better password
	if quality.Score < minScore {
		fmt.Println(RED + FAILURE + CLEAR +
			" Chosen password is too weak." +
			" Please select a better one.")
		return false, nil
	}

	// offer chance for a better password
	for choice != "y" && choice != "n" {
		if choice, err = verifyChoice(); err != nil {
			if err == liner.ErrPromptAborted {
				os.Exit(0)
			}
			return false, err
		}
	}

	switch choice {
	case "y":
		return true, nil
	case "n":
		return false, nil
	}
	return false, fmt.Errorf("Unreachable error reached")
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

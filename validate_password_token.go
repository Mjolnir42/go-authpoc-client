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

	"gopkg.in/resty.v5"
)

func validatePasswordToken(user, token string) error {
	var (
		err  error
		resp *resty.Response
	)

	if resp, err = resty.New().
		//SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetRootCertificate("server.pem").
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(3)).
		R().
		SetBasicAuth(user, token).
		Get(fmt.Sprintf(
			"%s/authenticate/validate",
			ApiAddress,
		)); err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return fmt.Errorf("Response code was: %s\n", resp.Status())
	}
	return nil
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

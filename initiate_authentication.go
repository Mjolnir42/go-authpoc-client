/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v5"
)

func initiateAuthentication(a *AuthRequest) error {
	var (
		err     error
		resp    *resty.Response
		payload []byte
	)

	if payload, err = json.Marshal(a); err != nil {
		return err
	}

	if resp, err = resty.New().
		SetRootCertificate("server.pem").
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(3)).
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(payload).
		Post(fmt.Sprintf(
			"%s/authenticate/",
			ApiAddress,
		)); err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return fmt.Errorf("Response code was: %s\n", resp.Status())
	}

	peerAuth := AuthRequest{}
	if err = json.Unmarshal(resp.Body(), &peerAuth); err != nil {
		return err
	}
	a.SetPeerKey(peerAuth.PublicKey())

	if err = a.SetRequest(peerAuth.Request.String()); err != nil {
		return err
	}
	return nil
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

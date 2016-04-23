/*-
* Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
* All rights reserved.
*
* Use of this source code is governed by a 2-clause BSD license
* that can be found in the LICENSE file.
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v5"
)

func serverRequest(url string, cipherText []byte, a *AuthRequest) (UserRequest, error) {
	var (
		cipherReply, bodyText []byte
		plainReply            *[]byte
		err                   error
		resp                  *resty.Response
		cipherLen             int
		uReq                  UserRequest
	)

	if resp, err = resty.New().
		SetRootCertificate("server.pem").
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(3)).
		R().
		SetHeader("Content-Type", "application/octet-stream").
		SetBody(cipherText).
		Post(url); err != nil {
		return uReq, err
	}

	if resp.StatusCode() >= 300 {
		return uReq, fmt.Errorf("Error: received %s\n", resp.Status())
	}

	bodyText = resp.Body()
	cipherReply = make([]byte, base64.StdEncoding.DecodedLen(len(bodyText)))
	if cipherLen, err = base64.StdEncoding.Decode(cipherReply, bodyText); err != nil {
		return uReq, err
	}

	if plainReply = decryptBytes(a, cipherReply[:cipherLen]); plainReply == nil {
		return uReq, fmt.Errorf("Failed to decrypt server reply")
	}

	if err = json.Unmarshal(*plainReply, &uReq); err != nil {
		return uReq, err
	}

	return uReq, nil
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix

//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gosendcloud.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gosendcloud

import (
	"bytes"
	"encoding/base64"
	"net/http"
)

const (
	transferProtocol = "https://"
	baseUrl          = "panel.sendcloud.sc/api"
	apiVersion       = "/v2"
)

// Config is to define config data
type Config struct {
	Path, Method string
	Body         []byte
}

// Request is to define the request data
type Request struct {
	PublicKey string
	SecretKey string
}

// Send is to send a new request
func (c Config) Send(r Request) (*http.Response, error) {

	// Set url
	url := transferProtocol + baseUrl + apiVersion + c.Path

	// Create basic authentication
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(r.PublicKey+":"+r.SecretKey))

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Define header
	request.Header.Set("Authorization", auth)
	request.Header.Set("Content-Type", "application/json")

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}

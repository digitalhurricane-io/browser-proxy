package main

import (
	"net"
	"net/http"
	"time"
)

func NewHTTPClientWithTimeout(seconds int) *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(seconds) * time.Second, // TCP connect timeout
		}).Dial,
		TLSHandshakeTimeout: time.Duration(seconds) * time.Second, // TLS handshake timeout
	}

	netClient := &http.Client{
		Timeout: time.Second * time.Duration(seconds), // response timeout
		Transport: netTransport,
	}

	return netClient
}

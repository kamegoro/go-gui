//go:build !release
// +build !release

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func init() {
	u, err := url.Parse("http://192.130.0.10:3200")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", httputil.NewSingleHostReverseProxy(u))
}
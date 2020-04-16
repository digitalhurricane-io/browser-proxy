package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)


func proxyHandler(w http.ResponseWriter, r *http.Request) {
	setCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	targetUrl := r.Header.Get("target")
	if targetUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := doReverseProxy(targetUrl, w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// reverse proxy to the given url
func doReverseProxy(target string, w http.ResponseWriter, r *http.Request) error {

	targetUrl, err := url.Parse(target)
	if err != nil {
		return err
	}

	proxy := NewSingleHostReverseProxy(targetUrl)

	proxy.ServeHTTP(w, r)

	return nil
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		req.URL.RawQuery = target.RawQuery

		req.Host = target.Host

		req.Header.Del("target")

		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	return &httputil.ReverseProxy{Director: director}
}

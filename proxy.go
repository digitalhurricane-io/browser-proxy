package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)


type OnwardData struct {
	DestinationUrl string `json:"url"`// full url eg. http://111.222.333.444:9000/my-service
	Data interface{} `json:"data"` // json data to be passed on
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	setCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var reqData OnwardData

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if reqData.DestinationUrl == "" {
		log.Println("No destination url provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var jsonForNextServer []byte
	if reqData.Data != nil {
		jsonForNextServer, err = json.Marshal(reqData.Data)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// put only the json that we want the next server to receive into the body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(jsonForNextServer))

	// we changed the body contents so we need to
	// change the content length to match
	r.ContentLength = int64(len(jsonForNextServer))

	err = doReverseProxy(reqData.DestinationUrl, w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// reverse proxy to the given url
func doReverseProxy(target string, w http.ResponseWriter, r *http.Request) error {

	targetUrl, err := url.Parse(target)
	if err != nil {
		return err
	}

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	proxy.ServeHTTP(w, r)

	return nil
}

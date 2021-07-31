package main

import (
	"log"
	"net/http"
)

type httpEntry func(http.ResponseWriter, *http.Request)

var httpEntryMap = map[string]httpEntry{
	apiV1Prefix:    apiV1Main,
	apiV1WebSocket: apiV1WebSocketHandler,
}

func setupHTTPHandler() {

	for key, val := range httpEntryMap {
		http.HandleFunc(key, val)
	}

	log.Printf("Server started at %s:%s\n", config.Host, config.Port)
	log.Fatal(http.ListenAndServe(config.Host+":"+config.Port, nil))
	//log.Fatal(http.ListenAndServeTLS(config.Host+":"+config.Port, config.TlsCert, config.TlsKey, nil))
}

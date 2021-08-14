package main

import (
	"net/http"
	"strings"
)

const apiV1Prefix = "/api1/"

func apiV1Main(w http.ResponseWriter, r *http.Request) {

}

func setupCrossOriginResponse(w *http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	requestedHeaders := r.Header.Get("Access-control-Request-Headers")
	method := r.Header.Get("Access-Control-Request-Method")
	(*w).Header().Set("Access-Control-Allow-Origin", origin) //for that specific origin
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", removeDupHeaderOptions("POST, GET, OPTIONS, PUT, DELETE, "+method))
	(*w).Header().Set("Access-Control-Allow-Headers", removeDupHeaderOptions("Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cookie, Biukop-Session, Biukop-Socket , "+requestedHeaders))
}

func removeDupHeaderOptions(inStr string) (out string) {
	headers := map[string]struct{}{}
	strings.ReplaceAll(inStr, " ", "")       // remove space
	headerArray := strings.Split(inStr, ",") // split
	for _, v := range headerArray {
		headers[v] = struct{}{} // same key will overwrite each other
	}
	out = ""
	for k, _ := range headers {
		if out != "" {
			out += ", "
		}
		out += k
	}
	return
}

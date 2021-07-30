package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type configStaticHtml struct {
	Dir         string
	StaticUrl   string
	StripPrefix string
}

type configuration struct {
	Host    string
	Port    string
	DSN     string
	TlsCert string
	TlsKey  string
	Static  []configStaticHtml
	Debug   bool
	TempDir string
	Session struct { //TODO: figure what is this intended for
		Guest bool
		Year  int //how many years
		Month int //how many years
		Day   int //how many years
	}
}

var configFile = "config.json"
var config = configuration{}

func (m *configuration) readConfig() (e error) {
	log.Printf("read Path config from %s", configFile)
	body, e := ioutil.ReadFile(configFile)
	if e != nil {
		log.Fatal("Cannot read config from " + configFile)
		return
	}
	e = json.Unmarshal(body, m)

	if config.Debug {
		log.Println(config)
	}
	return
}

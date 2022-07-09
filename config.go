package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type configStaticHtml struct {
	Dir         string
	StaticUrl   string
	StripPrefix string
	Sync        string
}

type configuration struct {
	Host     string
	Port     string
	DSN      string
	TlsCert  string
	TlsKey   string
	RSyncKey string
	Static   []configStaticHtml
	Debug    bool
	TempDir  string
	Session  struct { //TODO: figure what is this intended for
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

	// Check upload dir and defaults
	if !config.checkUploadDir() {
		log.Fatal("bad config file", configFile)
		return
	}

	if config.Debug {
		log.Println(config)
	}
	return
}

func (m *configuration) checkUploadDir() (valid bool) {
	valid = true
	for idx, node := range m.Static {
		if node.StaticUrl == "/" {
			if !fileExists(node.Dir) {
				valid = false
				log.Fatal(" html / not exist ", node)
			} else {
				// convert to absolute path : fileDir
				p, e := filepath.Abs(node.Dir)
				if e != nil {
					valid = false
					log.Fatal("bad html (webroot) dir ", node, e)
				}
				m.Static[idx].Dir = p + string(os.PathSeparator) //change it to absolute dir
			}
		}
	}

	// convert rsync key file to absolute dir
	p, e := filepath.Abs(config.RSyncKey)
	if e != nil {
		valid = false
		log.Fatal("bad html (webroot) dir ", config.RSyncKey, e)
	}
	m.RSyncKey = p //change it to absolute dir

	return
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return false
	}
	return true
}

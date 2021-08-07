package main

// [START import]
import (
	"log"
	"net/http"
)

// [END import]
// [START main_func]

func main() {

	err := config.readConfig() //
	if err != nil {
		log.Println(err)
		log.Fatalf("unable to read %s, program quit\n", configFile)
		return
	}

	// [START setting_port]
	port := config.Port
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}
	// [END setting_port]

	setupRootFileServer()

	//always last
	setupHTTPHandler()
}

// [END main_func]

func setupRootFileServer() {
	//root of doc
	for idx, node := range config.Static {
		log.Printf("setting up static %d with %+v\n", idx, node)
		fs := FileServerWith404(http.Dir(node.Dir), fileSystem404)
		http.Handle(node.StaticUrl, http.StripPrefix(node.StripPrefix, fs))
	}
}

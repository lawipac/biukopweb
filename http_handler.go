package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type httpEntry func(http.ResponseWriter, *http.Request)

var httpEntryMap = map[string]httpEntry{
	apiV1Prefix:    apiV1Main,
	apiV1WebSocket: apiV1WebSocketHandler,
	videoPrefix:    videoMain,
	formPrefix:     formMain,
}

func setupHTTPHandler() {

	for key, val := range httpEntryMap {
		http.HandleFunc(key, val)
	}

	log.Printf("Server started at %s:%s\n", config.Host, config.Port)
	log.Fatal(http.ListenAndServe(config.Host+":"+config.Port, nil))
	//log.Fatal(http.ListenAndServeTLS(config.Host+":"+config.Port, config.TlsCert, config.TlsKey, nil))
}

// FSHandler404 provides the function signature for passing to the FileServerWith404
type FSHandler404 = func(w http.ResponseWriter, r *http.Request) (doDefaultFileServe bool)

/*
FileServerWith404 wraps the http.FileServer checking to see if the url path exists first.
If the file fails to exist it calls the supplied FSHandle404 function
The implementation can choose to either modify the request, e.g. change the URL path and return true to have the
default FileServer handling to still take place, or return false to stop further processing, for example if you wanted
to write a custom response
e.g. redirects to root and continues the file serving handler chain
	func fileSystem404(w http.ResponseWriter, r *http.Request) (doDefaultFileServe bool) {
		//if not found redirect to /
		r.URL.Path = "/"
		return true
	}
Use the same as you would with a http.FileServer e.g.
	r.Handle("/", http.StripPrefix("/", mw.FileServerWith404(http.Dir("./staticDir"), fileSystem404)))
*/

func FileServerWith404(root http.FileSystem, handler404 FSHandler404) http.Handler {
	fs := http.FileServer(root)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//make sure the url path starts with /
		upath := r.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			r.URL.Path = upath
		}
		upath = path.Clean(upath)

		// attempt to open the file via the http.FileSystem
		f, err := root.Open(upath)

		if err != nil {
			if os.IsNotExist(err) {
				if handledBy404Handler(handler404, w, r) {
					return
				}
			}
		}

		// if successfully opened, check if it's dir
		defer func(f http.File) {
			err := f.Close()
			if err != nil {
				return //cannot close this file
			}
		}(f)

		//see if it's directory
		info, e := f.Stat()
		if e != nil {
			if handledBy404Handler(handler404, w, r) {
				return
			}
		}

		if info.IsDir() {
			//do we have index.html?
			f, err := root.Open(upath + "/index.html")
			if err != nil {
				//no we don't
				if handledBy404Handler(handler404, w, r) {
					return
				}
			} else {
				f.Close()
			}
		}
		// default serve
		fs.ServeHTTP(w, r)
	})
}

func fileSystem404(w http.ResponseWriter, r *http.Request) (doDefaultFileServe bool) {
	//if not found redirect to /
	// r.URL.Path = "/404.html" //not working as some directorys may not be feasible.
	// return true

	http.Redirect(w, r, "/404.html", http.StatusSeeOther)
	return false
}

//check if h is null, if not call this as handler.
func handledBy404Handler(h FSHandler404, w http.ResponseWriter, r *http.Request) bool {
	if h != nil {
		doDefault := h(w, r)
		if !doDefault {
			return true
		}
	}
	return false
}

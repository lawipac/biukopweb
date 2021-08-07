package main

import (
	"fmt"
	"net/http"
)

//display vimeo video in biukop brand

const videoPrefix = "/v/"

type vimeoPlayer struct {
	VideoId string

	playsinline int
	autoplay    int
	autopause   int
	loop        int
	background  int
	muted       int

	Title string
}

func videoMain(w http.ResponseWriter, r *http.Request) {
	videoVimeo(w, r)
}

func videoVimeo(w http.ResponseWriter, r *http.Request) {

	vimeo := getVimeoParams(r)

	pattern := `
<html>
<title> %s </title>
<body style="margin:0px">
    <style> iframe {width: 100vw; height: 100vh; overflow:hidden;} </style>
    <iframe src="%s" 
        scrolling="no"  frameborder="0"
        style="width: 100vw; height: 100vh; overflow:hidden;" 
        allow="autoplay; fullscreen" 
        webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
</body>

</html>`
	output := fmt.Sprintf(pattern, vimeo.Title, vimeo.getUrl())
	fmt.Fprintf(w, output)
}

func getVimeoParams(r *http.Request) (ret vimeoPlayer) {
	prefix := videoPrefix + "v/"
	ret.VideoId = r.URL.Path[len(prefix):]
	ret.Title = "Video"
	ret.autopause = 0
	ret.autoplay = 1
	ret.playsinline = 0
	ret.loop = 1
	ret.background = 0
	ret.muted = 0

	if r.URL.Path[:len(prefix)] == videoPrefix+"b/" {
		ret.playsinline = 1
		ret.background = 1
		ret.muted = 1 // autoplay video must be muted.
	}
	return
}

func (m *vimeoPlayer) getUrl() (ret string) {
	ret = fmt.Sprintf(
		"https://player.vimeo.com/video/%s?playsinline=%d&autoplay=%d&autopause=%d&loop=%d&background=%d&muted=%d",
		m.VideoId, m.playsinline, m.autoplay, m.autopause, m.loop, m.background, m.muted)
	return
}

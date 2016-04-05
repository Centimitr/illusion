package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
)

func DirectGet(url string) []byte {
	r, _ := http.Get(url)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

func PrerenderGet(url string) []byte {
	out, _ := exec.Command(PHANTOMJS_PATH, PHANTOMJS_SCRIPT, url, "0", "0", "desktop").Output()
	return out
}

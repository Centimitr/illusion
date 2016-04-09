// This package implements simple prerender using PhantomJS
package illusion

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
)

// Common http.get and directly return r.Body's content
func DirectGet(url string) []byte {
	r, _ := http.Get(url)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

func prerenderGet(url, viewport string) []byte {
	out, _ := exec.Command(PHANTOMJS_PATH, PHANTOMJS_SCRIPT, url, strconv.Itoa(PRERENDER_PROCESS_TIMEOUT), strconv.Itoa(PRERENDER_REFRESH_TIMEOUT), viewport).Output()
	return out
}

// Request and wait for page's loading use PhantomJS with default viewport.
func PrerenderGet(url string) []byte {
	return prerenderGet(url, PRERENDER_DEFAULT_VIEWPORT)
}

// Request and wait for page's loading use PhantomJS with desktop viewport.
func PrerenderGetDesktop(url string) []byte {
	return prerenderGet(url, "desktop")
}

// Request and wait for page's loading use PhantomJS with mobile viewport.
func PrerenderGetMobile(url string) []byte {
	return prerenderGet(url, "mobile")
}

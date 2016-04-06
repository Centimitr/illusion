package illusion

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

func prerenderGet(url, viewport string) []byte {
	out, _ := exec.Command(PHANTOMJS_PATH, PHANTOMJS_SCRIPT, url, PRERENDER_PROCESS_TIMEOUT, PRERENDER_REFRESH_TIMEOUT, viewport).Output()
	return out
}

func PrerenderGet(url string) []byte {
	return prerenderGet(url, PRERENDER_DEFAULT_VIEWPORT)
}

func PrerenderGetDesktop(url string) []byte {
	return prerenderGet(url, "desktop")
}

func PrerenderGetMobile(url string) []byte {
	return prerenderGet(url, "mobile")
}

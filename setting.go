package illusion

var (
	PRERENDER_PROCESS_TIMEOUT  = 5000
	PRERENDER_REFRESH_TIMEOUT  = 500
	PRERENDER_DEFAULT_VIEWPORT = "desktop"
)

func SetPrerenderProcessTimeout(t int) {
	PRERENDER_PROCESS_TIMEOUT = t
}

func SetPrerenderRefreshTimeout(t int) {
	PRERENDER_REFRESH_TIMEOUT = t
}

func SetPrerenderDefaultViewport(v string) {
	if v == "desktop" || v == "mobile" {
		PRERENDER_DEFAULT_VIEWPORT = v
	}
}

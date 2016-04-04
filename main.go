package main

import (
	// "fmt"
	"io/ioutil"
)

func main() {
	url := "http://docs.ngnice.com/api/ng/function/angular.equals"
	ioutil.WriteFile("direct.html", DirectGet(url), 0777)
	ioutil.WriteFile("prerender.html", PrerenderGet(url), 0777)
}

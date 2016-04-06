package main

import (
	// "fmt"
	"github.com/Centimitr/illusion"
	"io/ioutil"
)

func writeFile(url, directName, prerenderName string) {
	ioutil.WriteFile(directName, illusion.DirectGet(url), 0777)
	ioutil.WriteFile(prerenderName, illusion.PrerenderGet(url), 0777)
}

func main() {
	writeFile("http://docs.ngnice.com/api/ng/function/angular.equals", "direct.html", "prerender.html")
}

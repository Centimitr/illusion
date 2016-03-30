package main

import (
	"archive/zip"
	// "bytes"
	// "fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func exactZip(fp string) {
	r, _ := zip.OpenReader(fp)
	defer r.Close()
	for _, f := range r.File {
		rc, _ := f.Open()
		data, _ := ioutil.ReadAll(rc)
		dir, fn := filepath.Split(f.Name)
		if fn == "" {
			os.Mkdir(f.Name, 0777)
		} else {
			d, _ := os.Stat(dir)
			if d.IsDir() {
				os.Mkdir(dir, 0777)
			}
			ioutil.WriteFile(f.Name, data, 0777)
		}
	}
}

func checkPhantomJs(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func page(url string) []byte {
	out, _ := exec.Command(PHANTOMJS_PATH, PHANTOMJS_SCRIPT, url).Output()
	return out
}

func directGet(url string) []byte {
	r, _ := http.Get(url)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

var (
	PHANTOMJS_ZIP    = "phantomjs-2.1.1-windows.zip"
	PHANTOMJS_PATH   = "./phantomjs-2.1.1-windows/bin/phantomjs.exe"
	PHANTOMJS_SCRIPT = "visit.js"
)

func init() {
	if checkPhantomJs(PHANTOMJS_PATH) {
		exactZip(PHANTOMJS_ZIP)
	}
	if !checkPhantomJs(PHANTOMJS_PATH) {
		os.Exit(1)
	}
}

func main() {
	url := "http://docs.ngnice.com/api/ng/function/angular.equals"
	ioutil.WriteFile("1.html", page(url), 0777)
	ioutil.WriteFile("2.html", directGet(url), 0777)
}

package illusion

import (
	"archive/zip"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// ignore errors

func mustDownload(link, _ string) {
	r, _ := http.Get(link)
	defer r.Body.Close()
	d, _ := ioutil.ReadAll(r.Body)
	_, fn := filepath.Split(link)
	ioutil.WriteFile(fn, d, 0777)
}

func mustExactZip(fp string) {
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

func fileIsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func fileIsNotExist(path string) bool {
	return !fileIsExist(path)
}

// Now only windows support, you can easily modify settings for your environment
var (
	OS                 = "windows"
	PHANTOMJS_ZIP_LINK = "https://bitbucket.org/ariya/phantomjs/downloads/phantomjs-2.1.1-windows.zip"
	PHANTOMJS_ZIP_PATH = "phantomjs-2.1.1-windows.zip"
	PHANTOMJS_PATH     = "./phantomjs-2.1.1-windows/bin/phantomjs.exe"
	PHANTOMJS_SCRIPT   = "prerender.js"
)

func init() {
	if fileIsNotExist(PHANTOMJS_PATH) {
		if fileIsNotExist(PHANTOMJS_ZIP_PATH) {
			mustDownload(PHANTOMJS_ZIP_LINK, ".")
		}
		mustExactZip(PHANTOMJS_ZIP_PATH)
	}
}

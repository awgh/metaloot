package metaloot

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// Metaloot - another name for pirate treasure
func Metaloot(basedir string, uri string) {

	resp, err := http.Get(uri)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	u, err := url.Parse(uri)
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	filename := filepath.Join(basedir, u.Path)
	ioutil.WriteFile(filename, b, os.FileMode(int(0600)))
	if err != nil {
		log.Println(err.Error())
		return
	}

	r := bufio.NewReader(bytes.NewReader(b))
	more := true
	for more {
		var line []byte
		line, more, err = r.ReadLine()
		if err != nil {
			log.Println(err.Error())
			return
		}
		Metaloot(basedir, uri+"/"+string(line))
	}
}

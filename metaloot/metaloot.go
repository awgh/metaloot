package metaloot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func mkdirP(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

// Metaloot - another name for pirate treasure
func Metaloot(basedir string, uri string) error {
	log.Println("Metaloot called: ", basedir, uri)
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("%d: %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {

		if line == "" || strings.Contains(line, " ") {
			continue // spaces cause a resource unavailable error
		}

		if err := Metaloot(basedir, uri+"/"+line); err != nil {
			log.Println(err.Error())
			log.Println("PATH FRAGMENT: ", u.Path)
			filename := filepath.Join(basedir, u.Path)
			mkdirP(filename)
			if err := ioutil.WriteFile(filename, b, os.FileMode(int(0600))); err != nil {
				log.Println(err.Error())
				return err
			}
		}
	}
	log.Println("No more lines in", uri)
	return nil
}

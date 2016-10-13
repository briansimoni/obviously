package main

import (
	"io/ioutil"
	"os"
	"bufio"
	"regexp"
)

func main() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		f, err:= os.Open(f.Name())
		defer f.Close()
		if err != nil {
			continue
		}
		s := bufio.NewScanner(f)
		fixed := ""
		for s.Scan() {
			l := s.Text()
			r := regexp.MustCompile(`\s\s\s\s`)
			tabs := r.ReplaceAllString(l, "\t")
			fixed += tabs + "\n"
		}
		err = ioutil.WriteFile(f.Name(), []byte(fixed), 775)

	}
}

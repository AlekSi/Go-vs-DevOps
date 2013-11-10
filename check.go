package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("go-vs-devops")
	if err != nil {
		log.Fatal(err)
	}

	used := make(map[string]bool)
	for _, f := range files {
		used["go-vs-devops/"+f.Name()] = false
	}

	f, err := os.Open("go-vs-devops.slide")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		parts := strings.Split(strings.TrimSpace(line), " ")
		if len(parts) > 1 && ((parts[0] == ".image") || (parts[0] == ".code") || (parts[0] == ".play")) {
			_, ok := used[parts[1]]
			if !ok {
				log.Printf("%s not found", parts[1])
			}
			used[parts[1]] = true
		}
	}

	for f, u := range used {
		if !u {
			log.Printf("%s not used", f)
		}
	}
}

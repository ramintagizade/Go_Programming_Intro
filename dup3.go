package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, " %v ", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, cnt := range counts {
		fmt.Println("%s %d", line, cnt)
	}
}

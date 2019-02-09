package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main () {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, cnt := range counts {
		fmt.Printf("%s %d " , line , cnt)
	}
}
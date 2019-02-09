package main 

import (
	"fmt"
	"os"
	"bufio"
)



func main() {

	files := os.Args[1:]
	for _, file :=range files {
		f, err:= os.Open(file)
		if err!=nil {
			fmt.Errorf("Error")
		} 
		readFile(f)
	}

	
}

func readFile(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}
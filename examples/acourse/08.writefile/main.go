package main

import "fmt"
import "os"

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	f.WriteString("Hello!")
}

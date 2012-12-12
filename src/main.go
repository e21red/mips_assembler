package main

import (
//	"assembler"
	"fmt"
	"ioutil"
	"os"
	"strings"
)

func main() {
	content, err := iotuil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error Reading File:", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	/* Assemble, get symbols for header */
	
	/* Print header */

	/* Print bytes  */

}
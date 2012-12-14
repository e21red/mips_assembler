package main

import (
	"assembler"
	"fmt"
	"header"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	D := "11110000111100001111000011110000"

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error Reading File:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	
	/* Assemble, get symbols for header */
		machineCode, symbolTable, cmds := assembler.Assemble(lines)
	
	/* Print header */
	dataCode := header.Data(symbolTable)
	headerCode := header.Header(symbolTable, cmds)
	
	var output string

	output += D
	output += headerCode

	for _, entry := range dataCode {
		x := "" + entry
		output += x
	}
	/* Print data section */
	output += D

	/* Print bytes  */
	for _, command := range machineCode {
		output += command
	}
	output += D

	ioutil.WriteFile("asm.out", []byte(output), 0666)
}
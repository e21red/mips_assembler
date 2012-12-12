package assembler

import (
	"opcodes"
	"regexp"
	"strings"
	)

func RemoveComments(input []string) (output []string) {
	re := regexp.MustCompile("[\".*\"]*;.+")
	for _, line := range input {
		line = string(re.ReplaceAll([]byte(line), []byte(" ")))
		output = append(output, line)
	}
	return output
}

func removeWhitespace(input []string) (output []string) {
	for _, line := range input {
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "\t", "", -1)
		output = append(output, line)
	}
	return output
}

func Assemble(input []string) (output []string) {
	OPCODES := opcodes.OPCODES()

	input = RemoveComments(input)
	for _, line := range input {
		/* Parse, then map opcodes to commands */
			words = strings.Split(strings.TrimSpace(line))
		
		output = append(output, line)
	}
	return output
}
package assembler

import (
	"fmt"
	"opcodes"
	"regexp"
	"strconv"
	"strings"
	)

func RemoveComments(input []string) (output []string) {
	re := regexp.MustCompile("[\".*\"]*;.+")
	for _, line := range input {
		output = append(output, string(re.ReplaceAll([]byte(line), []byte(" "))))
	}
	return output
}

func Assemble(input []string) (output []string) {
	OPCODES := opcodes.OPCODES()
	REGISTERS := opcodes.REGISTERS()

	input = RemoveComments(input)
	for linenum, line := range input {
		var bytecode string
		
		line = strings.Replace(line, ",", "", -1)
		fields := strings.Fields(line)		
		
		op, fn, ft := codename(fields[0], OPCODES, linenum)
		switch ft {
		case 0:
			// No-op, all zeroes
			bytecode = fmt.Sprintf("%32b", 0)
		case 1:		
			// R-type, arithmetic ops
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[3], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), 0, fn)
		case 2:
			// R-type, mult and div
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[1], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), 0, 0, fn)
		case 3:
			// R-type, shift without variable
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, 0, regname(fields[2], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), fields[3], fn)
		case 4:
			// R-type, shift with variable
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[3], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), 0, fn)
		case 5:
			// R-type, JR
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[1], REGISTERS, linenum), 0, 0, 0, fn)
		case 6:
			// I-type arithmetic ops
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[2], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), twosComplement(fields[3]))
		case 7:
			// I-type, be, bne type 
				bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[1], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), twosComplement(fields[3])) // addr+pc ?
		case 8:
			// I-type, BGEZ only
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[1], REGISTERS, linenum), 1, twosComplement(fields[2]))
		case 9:
			// I-type, non-BGEZ
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[1], REGISTERS, linenum), 0, twosComplement(fields[2]))
		case 10:
			// I-type, using parenthesized offsets
			items := strings.Split(fields[2], "(")
			ofs := strings.Trim(items[2], ")")
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(items[1], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), twosComplement(ofs))
		case 11:
			// I-type LUI
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, 0, regname(fields[1], REGISTERS, linenum), twosComplement(fields[2]))
		case 12:
			// Jumps
			bytecode = fmt.Sprintf("%06b%026b", op, twosComplement(fields[1]))
		case 13:
			// Move from hi/lo
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, 0, 0, regname(fields[1], REGISTERS, linenum), 0, fn)
		case 14:
			bytecode = fmt.Sprintf("%06b%020b%06b", op, 0, fn)
		}
		

		output = append(output, bytecode)
	}
	return output
}

func codename(code string, OPCODES map[string]opcodes.Byterep, linenum int) (int, int, int) {	
	_, exists := OPCODES[code]
	if !exists {
		fmt.Println("Error: command", code, "not found, line", linenum)
//		return
	}

	bytes := OPCODES[code]
	return bytes.InstructionCode, bytes.FunctionCode, bytes.Format
}

func regname(reg string, REGISTERS map[string]int, linenum int) (regnum int) {
	_, exists := REGISTERS[reg]
	if !exists {
		fmt.Println("Error: register", reg, "not found, line", linenum)
//		return
	}

	return  REGISTERS[reg]
}

func twosComplement(s string) uint16 {
	i, _ := strconv.Atoi(s)
	return uint16(i)
}
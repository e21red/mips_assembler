package assembler

import (
	"fmt"
	"header"
	"opcodes"
	"regexp"
	"strconv"
	"strings"
	)

/*
 * Strips off everything from a line after a semicolon. 
 * Woe to you if you have a semicolon in a string and don't escape it.
 * Also gets rid of blank lines
 */
func removeComments(input []string) (output []string) {
	re := regexp.MustCompile("[^\\\\];.+")
	for _, line := range input {
		if line != "" {
			output = append(output, string(re.ReplaceAll([]byte(line), []byte(" "))))
		}
	}
	return output
}

/*
 * Take in the lines of assembly code and produce the corresponding machine code. Expects the machine code to follow a
 * specific format, which is detailed in the readme. Does not handle pseudo-ops. Returns an array of individual machine instructions.
 */
func Assemble(input []string) (output []string, symbols map[string]header.Entry, cmds int) {
	WORDSIZE := 4
	OPCODES := opcodes.OPCODES()
	REGISTERS := opcodes.REGISTERS()
	
	input = removeComments(input)
	symbols = make(map[string]header.Entry)
	// first pass for loop to grab .data, .asciiz
	var (
		symbol string
		lc uint16
		data bool
	)
	
	// First pass to find all our symbols
	for linenum, line := range input {
		fields := strings.Fields(line)
		if data {
			/* Now, we require all variable declarations to be of the form name: .type 16 */
			for i, field := range fields {
				if field[0] == '.' {
					symbol = strings.TrimRight(fields[i-1], ":")
					oldlc := lc
					if field == ".space" { 
						conv, _ := strconv.Atoi(fields[i+1])
						lc += uint16(conv)
					} else if field == ".word" {
						lc +=  uint16(WORDSIZE * len(strings.Split(fields[i+1], ",")))
					} else if field == ".asciiz" {
						fields[i+1] = strings.Trim(fields[i+1], "\"")
						lc += uint16(len(fields[i+1]) + 1)
					} else {
						fmt.Println("Invalid instruction", field, "at line",linenum)
						break
					}
					symbols[symbol] = header.Entry{uint32(lc), uint32(lc - oldlc), []byte(fields[i+1]), true} // Instruction address
				}
			}
		} else {
			if strings.Contains(fields[0], ":") {
				symbols[symbol] = header.Entry{uint32(lc), 4, []byte("0000"), false} // Nobody cares what the data is, just better be 32 long, so 4-byte array
			}
			cmds++
		}
		if data && strings.Contains(line, ".text") {
			cmds--
			data = false
		}
		if !data && strings.Contains(line, ".data") {
			cmds--
			data = true
		}
		lc += 4
	}

	data = false
	// second pass to replace and assemble
	for linenum, line := range input {
		var bytecode string
		
		line = strings.Replace(line, ",", "", -1)
		fields := strings.Fields(line)		
		fields[0] = strings.ToUpper(fields[0])

		op, fn, ft := codename(fields[0], OPCODES, linenum)
		if op == -1 {
			break
		}
		switch ft {
		case 0:
			// No-op, all zeroes
			bytecode = fmt.Sprintf("%032b", 0)
		case 1:		
			// R-type, arithmetic ops
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[3], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), 0, fn)
		case 2:
			// R-type, mult and div
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[1], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), 0, 0, fn)
		case 3:
			// R-type, shift without variable
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, 0, regname(fields[2], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), immediate(fields[3], symbols), fn)
		case 4:
			// R-type, shift with variable
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[3], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), 0, fn)
		case 5:
			// R-type, JR
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, regname(fields[1], REGISTERS, linenum), 0, 0, 0, fn)
		case 6:
			// I-type arithmetic ops
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[2], REGISTERS, linenum), regname(fields[1], REGISTERS, linenum), immediate(fields[3], symbols))
		case 7:
			// I-type, be, bne type 
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[1], REGISTERS, linenum), regname(fields[2], REGISTERS, linenum), immediate(fields[3], symbols))
		case 8:
			// I-type, BGEZ only
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[1], REGISTERS, linenum), 1, immediate(fields[2], symbols))
		case 9:
			// I-type, non-BGEZ
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, regname(fields[1], REGISTERS, linenum), 0, immediate(fields[2], symbols))
		case 10:
			// I-type, using parenthesized offsets
			items := strings.Split(fields[2], ")")
			ofs := strings.Trim(items[0], "(")
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, immediate(items[1], symbols), regname(fields[1], REGISTERS, linenum), immediate(ofs, symbols))
		case 11:
			// I-type LUI
			bytecode = fmt.Sprintf("%06b%05b%05b%016b", op, 0, regname(fields[1], REGISTERS, linenum), immediate(fields[2], symbols))
		case 12:
			// Jumps
			bytecode = fmt.Sprintf("%06b%026b", op, immediate(fields[1], symbols))
		case 13:
			// Move from hi/lo
			bytecode = fmt.Sprintf("%06b%05b%05b%05b%05b%06b", op, 0, 0, regname(fields[1], REGISTERS, linenum), 0, fn)
		case 14:
			// Syscall
			bytecode = fmt.Sprintf("%06b%020b%06b", op, 0, fn)
		default:
			break
		}
		output = append(output, bytecode)
	}
	return output, symbols, cmds
}

/*
 * Takes an operation and a map of opcodes to their representations. Returns the individual parts of that 
 * representation after error-checking.
 */
func codename(code string, OPCODES map[string]opcodes.Byterep, linenum int) (int, int, int) {	
	_, exists := OPCODES[code]
	if !exists {
		return -1,-1,-1
	}

	bytes := OPCODES[code]
	return bytes.InstructionCode, bytes.FunctionCode, bytes.Format
}

/*
 * Takes a register and a map of register names to their numbers, returns the number after error-checking
 */
func regname(reg string, REGISTERS map[string]int, linenum int) (regnum int) {
	_, exists := REGISTERS[reg]
	if !exists {
		fmt.Println("Error: register", reg, "not found, line", linenum)
	}

	return  REGISTERS[reg]
}

/*
 * Converts a string containing an int to a uint16. Why? Because fmt.Printf("%b", -2) looks like -000000010.
 */
func twosComplement(s string) uint16 {
	i, _ := strconv.Atoi(s)
	return uint16(i)
}

/*
 * Decides whether of not to treat an immediate value in an I-type operation as a memory ref or a constant
 */
func immediate(s string, symbols map[string]header.Entry) uint16 {
	_, ok := symbols[s]
	if ok {
		return uint16(symbols[s].Location) // Dangerous looking, but will work because it was 16 bits to start with
	}
	return twosComplement(s)
}
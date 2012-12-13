package opcodes

import (
	"fmt"
//	"regexp"
)
type Byterep struct {
	InstructionCode int
	FunctionCode int
	Format int
}

func OPCODES() map[string]Byterep {
	CODES := make(map[string]Byterep)
	
	CODES["ADD"] 	= Byterep{0, 32,1}
	CODES["ADDI"] 	= Byterep{8, 0, 6}
	CODES["ADDIU"] 	= Byterep{9, 0, 6}
	CODES["ADDU"] 	= Byterep{0, 33, 1}
	CODES["AND"] 	= Byterep{0, 36, 1}
	CODES["ANDI"] 	= Byterep{12, 0, 6}
	CODES["BEQ"] 	= Byterep{4, 0, 7}
 	CODES["BGEZ"] 	= Byterep{1, 0, 8} 
	//CODES["BGEZAL"] =Byterep{9, 0}
	CODES["BGTZ"] 	= Byterep{7, 0, 9}
	CODES["BLEZ"] 	= Byterep{6, 0, 9}
	CODES["BLTZ"] 	= Byterep{1, 0, 9}
	//CODES["BLTZAL"] = Byterep{9, 0}
	CODES["BNE"] 	= Byterep{5, 0, 7}
	CODES["DIV"] 	= Byterep{0, 26, 2} 
	CODES["DIVU"] 	= Byterep{0, 27, 2} 
	CODES["J"] 	= Byterep{2, 0, 12}
	CODES["JAL"] 	= Byterep{3, 0, 12}
	CODES["JR"] 	= Byterep{0, 8, 5}	
	CODES["LB"] 	= Byterep{32, 0, 10}
	CODES["LUI"] 	= Byterep{15, 0, 11}
	CODES["LW"] 	= Byterep{35, 0, 10}
	CODES["MFHI"] 	= Byterep{0, 16, 13} 
	CODES["MFLO"] 	= Byterep{0, 18, 13} 
	CODES["MULT"] 	= Byterep{0, 24, 2}
	CODES["MULTU"] 	= Byterep{0, 25, 2}
	CODES["NOOP"] 	= Byterep{0, 0, 0}
	CODES["OR"] 	= Byterep{0, 37, 1}
	CODES["ORI"] 	= Byterep{0, 13, 6}
	CODES["SB"] 	= Byterep{40, 0, 10}
	CODES["SLL"] 	= Byterep{0, 0, 3}
	CODES["SLLV"] 	= Byterep{0, 4, 4}
	CODES["SLT"] 	= Byterep{0, 42, 1}
	CODES["SLTI"] 	= Byterep{10, 0, 6}
	CODES["SLTIU"] 	= Byterep{11, 0, 6}
	CODES["SLTU"] 	= Byterep{0, 43, 1}
	CODES["SRA"] 	= Byterep{0, 3, 3}
	CODES["SRL"] 	= Byterep{0, 2, 3}
	CODES["SRLV"] 	= Byterep{0, 6, 4}
	CODES["SUB"] 	= Byterep{0, 34, 1}
	CODES["SUBU"] 	= Byterep{0, 35, 1}
	CODES["SW"] 	= Byterep{43, 0, 10}
	CODES["SYSCALL"]= Byterep{0, 12, 14}
	CODES["XOR"] 	= Byterep{0, 38, 1}
	CODES["XORI"] 	= Byterep{14, 0, 6}

	return CODES
}

func REGISTERS() map[string]int {
	REGS := make(map[string]int)
	/*
	re := regexp.MustCompile("\$\d{1,2}")
	if re.match() {

	}*/

	REGS["$zero"] = 0
	REGS["$at"] = 1
	REGS["$v0"] = 2
	REGS["$v1"] = 3
	REGS["$a0"] = 4
	REGS["$a1"] = 5
	REGS["$a2"] = 6
	REGS["$a3"] = 7
	REGS["$t0"] = 8
	REGS["$t1"] = 9
	REGS["$t2"] = 10
	REGS["$t3"] = 11
	REGS["$t4"] = 12
	REGS["$t5"] = 13
	REGS["$t6"] = 14
	REGS["$t7"] = 15
	REGS["$t8"] = 24
	REGS["$t9"] = 25
	REGS["$s0"] = 16
	REGS["$s1"] = 17
	REGS["$s2"] = 18
	REGS["$s3"] = 19
	REGS["$s4"] = 20
	REGS["$s5"] = 21
	REGS["$s6"] = 22
	REGS["$s7"] = 23
	REGS["$s8"] = 30
	REGS["$k0"] = 26
	REGS["$k1"] = 27
	REGS["$gp"] = 28
	REGS["$sp"] = 29
	REGS["$ra"] = 31

	return REGS
}

func Int2Byte(n int) []byte {
	return []byte(fmt.Sprintf("%06b", n))
}


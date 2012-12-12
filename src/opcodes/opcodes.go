package opcodes

import "fmt"

type opcode string
type byterep struct {
	instructionCode int
	functionCode int
}

func OPCODES() map[opcode]byterep {
	var CODES map[opcode]byterep
	
	CODES["ADD"] 	= byterep{0, 32}
	CODES["ADDI"] 	= byterep{8, 0}
	CODES["ADDIU"] 	= byterep{9, 0}
	CODES["ADDU"] 	= byterep{0, 33}
	CODES["AND"] 	= byterep{0, 36}
	CODES["ANDI"] 	= byterep{12, 0}
	CODES["BEQ"] 	= byterep{4, 0}
	CODES["BGEZ"] 	= byterep{1, 0}
	//CODES["BGEZAL"] =byterep{9, 0}
	CODES["BGTZ"] 	= byterep{7, 0}
	CODES["BLEZ"] 	= byterep{6, 0}
	CODES["BLTZ"] 	= byterep{1, 0}
	//CODES["BLTZAL"] = byterep{9, 0}
	CODES["BNE"] 	= byterep{5, 0}
	CODES["DIV"] 	= byterep{0, 26}
	CODES["DIVU"] 	= byterep{0, 27}
	CODES["J"] 	= byterep{2, 0}
	CODES["JAL"] 	= byterep{3, 0}
	CODES["JR"] 	= byterep{0, 8}
	CODES["LB"] 	= byterep{32, 0}
	CODES["LUI"] 	= byterep{15, 0}
	CODES["LW"] 	= byterep{35, 0}
	CODES["MFHI"] 	= byterep{0, 16}
	CODES["MFLO"] 	= byterep{0, 18}
	CODES["MULT"] 	= byterep{0, 24}
	CODES["MULTU"] 	= byterep{0, 25}
	CODES["NOOP"] 	= byterep{0, 0}
	CODES["OR"] 	= byterep{0, 37}
	CODES["ORI"] 	= byterep{0, 13}
	CODES["SB"] 	= byterep{40, 0}
	CODES["SLL"] 	= byterep{0, 0}
	CODES["SLV"] 	= byterep{0, 4}
	CODES["SLT"] 	= byterep{0, 42}
	CODES["SLTI"] 	= byterep{10, 0}
	CODES["SLTIU"] 	= byterep{11, 0}
	CODES["SLTU"] 	= byterep{0, 43}
	CODES["SRA"] 	= byterep{0, 3}
	CODES["SRL"] 	= byterep{0, 2}
	CODES["SRLV"] 	= byterep{0, 6}
	CODES["SUB"] 	= byterep{0, 34}
	CODES["SUBU"] 	= byterep{0, 35}
	CODES["SW"] 	= byterep{43, 0}
	CODES["SYSCALL"]= byterep{0, 12}
	CODES["XOR"] 	= byterep{0, 38}
	CODES["XORI"] 	= byterep{14, 0}

	return CODES
}


func Int2Byte(n int) string {
	return fmt.Sprintf("%b", n)
}


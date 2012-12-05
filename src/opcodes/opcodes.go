package opcodes

type opcode string
type byterep struct {
	instructionCode int
	functionCode int
}

var CODES map[opcode]byterep

CODES["ADD"] 	= (0, 32)
CODES["ADDI"] 	= (8, 0)
CODES["ADDIU"] 	= (9, 0)
CODES["ADDU"] 	= (0, 33)
CODES["AND"] 	= (0, 36)
CODES["ANDI"] 	= (12, 0)
CODES["BEQ"] 	= (4, 0)
CODES["BGEZ"] 	= (1, 0)
//CODES["BGEZAL"] = (9, 0)
CODES["BGTZ"] 	= (7, 0)
CODES["BLEZ"] 	= (6, 0)
CODES["BLTZ"] 	= (1, 0)
//CODES["BLTZAL"] = (9, 0)
CODES["BNE"] 	= (5, 0)
CODES["DIV"] 	= (0, 26)
CODES["DIVU"] 	= (0, 27)
CODES["J"] 	= (2, 0)
CODES["JAL"] 	= (3, 0)
CODES["JR"] 	= (0, 8)
CODES["LB"] 	= (32, 0)
CODES["LUI"] 	= (15, 0)
CODES["LW"] 	= (35, 0)
CODES["MFHI"] 	= (0, 16)
CODES["MFLO"] 	= (0, 18)
CODES["MULT"] 	= (0, 24)
CODES["MULTU"] 	= (0, 25)
CODES["NOOP"] 	= (0, 0)
CODES["OR"] 	= (0, 37)
CODES["ORI"] 	= (0, 13)
CODES["SB"] 	= (40, 0)
CODES["SLL"] 	= (0, 0)
CODES["SLV"] 	= (0, 4)
CODES["SLT"] 	= (0, 42)
CODES["SLTI"] 	= (10, 0)
CODES["SLTIU"] 	= (11, 0)
CODES["SLTU"] 	= (0, 43)
CODES["SRA"] 	= (0, 3)
CODES["SRL"] 	= (0, 2)
CODES["SRLV"] 	= (0, 6)
CODES["SUB"] 	= (0, 34)
CODES["SUBU"] 	= (0, 35)
CODES["SW"] 	= (43, 0)
CODES["SYSCALL"]= (0, 12)
CODES["XOR"] 	= (0, 38)
CODES["XORI"] 	= (14, 0)

func int2Byte(n int) string {
	return sprintf("%b"
}
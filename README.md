mips_assembler
==============

Golang implementation of a mips assembler, by Carl Butt, Abhinav Shrestha, Shankara Pailoor

Expects assembly in a certain format:

1) No pseduo-instructions permitted.
2) All data declarations must be one line. e.g. myNum: .word 2
3) Only .asciiz, .word, and .space are supported as data types.
4) Requires that semicolons in .asciiz definitions by escaped by a backslash, will treat them as comments otherwise.

Accepts the name of an input file, will assemble to the binary file asm.out.

When trying to compile, make sure that $GOPATH is set. The script setup.sh will do that for you.

This program has been tested only with the 6g compiler.
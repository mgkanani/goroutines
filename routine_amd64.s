// +build amd64, gc

#include "go_asm.h"
#include "textflag.h"

// func goRoutine() *g
TEXT ·goRoutine(SB),NOSPLIT,$0-8
	MOVQ (TLS), R13
	MOVQ R13, ret+0(FP)
	RET

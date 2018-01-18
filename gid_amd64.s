#include "go_asm.h"
#include "textflag.h"

// func getgid() int64
TEXT ·getgid(SB),$0
	MOVQ	TLS, CX 	// get goroutine
	MOVQ	g_goid(CX), AX	// get id from g
	MOVQ	AX, ret+0(FP)
	RET

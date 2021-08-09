 88         0x0199 00409 (f1.go:21) MOVQ    CX, (SP)
 89         0x019d 00413 (f1.go:21) MOVQ    $1, 8(SP)
 90         0x01a6 00422 (f1.go:21) MOVQ    $1, 16(SP)
 91         0x01af 00431 (f1.go:21) CALL    log.Println(SB)
 92         0x01b4 00436 (f1.go:22) MOVQ    $120000000000, AX
 93         0x01be 00446 (f1.go:22) MOVQ    AX, (SP)
 94         0x01c2 00450 (f1.go:22) CALL    time.Sleep(SB)
 95         0x01c7 00455 (f1.go:26) XORPS   X0, X0
 96         0x01ca 00458 (f1.go:26) MOVUPS  X0, ""..autotmp_8+104(SP)
 97         0x01cf 00463 (f1.go:26) MOVUPS  X0, ""..autotmp_8+120(SP)
 98         0x01d4 00468 (f1.go:26) LEAQ    type.string(SB), AX
 99         0x01db 00475 (f1.go:26) MOVQ    AX, ""..autotmp_8+104(SP)
100         0x01e0 00480 (f1.go:26) LEAQ    ""..stmp_2(SB), AX
101         0x01e7 00487 (f1.go:26) MOVQ    AX, ""..autotmp_8+112(SP)
102         0x01ec 00492 (f1.go:26) MOVQ    "".data.len+72(SP), AX
103         0x01f1 00497 (f1.go:26) MOVQ    AX, (SP)
104         0x01f5 00501 (f1.go:26) PCDATA  $1, $2
105         0x01f5 00501 (f1.go:26) CALL    runtime.convT64(SB)
106         0x01fa 00506 (f1.go:26) MOVQ    8(SP), AX
107         0x01ff 00511 (f1.go:26) LEAQ    type.int(SB), CX
108         0x0206 00518 (f1.go:26) MOVQ    CX, ""..autotmp_8+120(SP)
109         0x020b 00523 (f1.go:26) MOVQ    AX, ""..autotmp_8+128(SP)
110         0x0213 00531 (f1.go:26) LEAQ    ""..autotmp_8+104(SP), AX
111         0x0218 00536 (f1.go:26) MOVQ    AX, (SP)
112         0x021c 00540 (f1.go:26) MOVQ    $2, 8(SP)
113         0x0225 00549 (f1.go:26) MOVQ    $2, 16(SP)
114         0x022e 00558 (f1.go:26) PCDATA  $1, $0
115         0x022e 00558 (f1.go:26) CALL    log.Println(SB)

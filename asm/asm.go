package asm

import (
	"bytes"
	"fmt"
	"os"
)

type Flag string

const (
	EQ Flag = "EQ"
	NE Flag = "NE"
	CS Flag = "CS"
	HS Flag = "HS"
	CC Flag = "CC"
	LO Flag = "LO"
	MI Flag = "MI"
	PL Flag = "PL"
	VS Flag = "VS"
	VC Flag = "VC"
	HI Flag = "HI"
	LS Flag = "LS"
	GE Flag = "GE"
	LT Flag = "LT"
	GT Flag = "GT"
	LE Flag = "LE"
	AL Flag = "AL"
	NI Flag = ""
)

type AssemblyWriter struct {
	BufferIndex int
	Buffers []*RegionBuffer
}

type RegionBuffer struct {
	Region int
	Buffer bytes.Buffer
}

func NewAssemblyWriter() *AssemblyWriter {
	return &AssemblyWriter{
		BufferIndex: 0,
		Buffers: []*RegionBuffer{
			{
				Region: 0,
				Buffer: bytes.Buffer{},
			},
		},
	}
}

func (w *AssemblyWriter) NewRegion(region int)  {
	w.BufferIndex += 1
	w.Buffers = append(w.Buffers, &RegionBuffer{
		Region: region,
		Buffer: bytes.Buffer{},
	})
}

func (w *AssemblyWriter) ExitRegion()  {
	w.BufferIndex -= 1
}

func (w *AssemblyWriter) WriteToFile(filename string) {
	f, err := os.Create(fmt.Sprintf("output/%s.s", filename))

	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, regBuffer := range w.Buffers {
		f.Write(regBuffer.Buffer.Bytes())
		f.WriteString("\n")
	}
}

func (w *AssemblyWriter) Raw(s string) {
	regionBuffer := w.Buffers[w.BufferIndex]
	regionBuffer.Buffer.WriteString(s)
}

func (w *AssemblyWriter) Label(label string) {
	w.Raw(label + ":\n")
}

func (w *AssemblyWriter) SkipLine() {
	w.Raw("\n")
}

func (w *AssemblyWriter) Add(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tADD%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Adc(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tADC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Sub(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tSUB%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Sbc(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tSBC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Rsb(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tRSB%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Rsc(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tRSC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Mov(dst string, val string, flag Flag) {
	instr := fmt.Sprintf("\tMOV%s %s, %s\n", flag, dst, val)

	w.Raw(instr)
}

func (w *AssemblyWriter) Mvn(dst string, val string, flag Flag) {
	instr := fmt.Sprintf("\tMVN%s %s, %s\n", flag, dst, val)

	w.Raw(instr)
}

func (w *AssemblyWriter) And(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tAND%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Eor(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tEOR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Lsl(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tLSL%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Lsr(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tLSR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Asr(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tASR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Ror(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tROR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Orr(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tORR%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Bic(dst string, val1 string, val2 string, flag Flag) {
	instr := fmt.Sprintf("\tBIC%s %s, %s, %s\n", flag, dst, val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Cmp(val1 string, val2 string) {
	instr := fmt.Sprintf("\tCMP %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Cmn(val1 string, val2 string) {
	instr := fmt.Sprintf("\tCMN %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Tst(val1 string, val2 string) {
	instr := fmt.Sprintf("\tTST %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Teq(val1 string, val2 string) {
	instr := fmt.Sprintf("\tTEQ %s, %s\n", val1, val2)

	w.Raw(instr)
}

func (w *AssemblyWriter) Ldr(dst string, addr string, flag Flag, offset int) {
	instr := fmt.Sprintf("\tLDR%s %s, [%s, #%d]\n", flag, dst, addr, offset)

	w.Raw(instr)
}

func (w *AssemblyWriter) Str(dst string, addr string, flag Flag, offset int) {
	instr := fmt.Sprintf("\tSTR%s %s, [%s, #%d]\n", flag, dst, addr, offset)

	w.Raw(instr)
}

func (w *AssemblyWriter) Ldmfd(src string, regs []string) {
	regsStr := ""

	for i, reg := range regs {
		regsStr += reg

		if i != len(regs) - 1 {
			regsStr += ", "
		}
	}


	instr := fmt.Sprintf("\tLDMFD %s, !{%s}\n", src, regsStr)

	w.Raw(instr)
}

func (w *AssemblyWriter) Stmfd(dst string, regs []string) {
	regsStr := ""

	for i, reg := range regs {
		regsStr += reg

		if i != len(regs) - 1 {
			regsStr += ", "
		}
	}


	instr := fmt.Sprintf("\tSTMFD %s, !{%s}\n", dst, regsStr)

	w.Raw(instr)
}

func (w *AssemblyWriter) B(label string, flag Flag) {
	instr := fmt.Sprintf("\tB%s %s\n", flag, label)

	w.Raw(instr)
}

func (w *AssemblyWriter) Bl(label string, flag Flag) {
	instr := fmt.Sprintf("\tBL%s %s\n", flag, label)

	w.Raw(instr)
}

func (w *AssemblyWriter) Fill(addr string, value string) {
	instr := fmt.Sprintf("\t%s FILL %s\n", addr, value)

	w.Raw(instr)
}

func (w *AssemblyWriter) End() {
	w.Raw("\tEND\n")
}

func (w *AssemblyWriter) Comment(comment string, tabs int) {
	instr := ""

	for i := 0; i < tabs; i++ {
		instr += "\t"
	}

	instr = fmt.Sprintf("%s/* %s */\n", instr, comment)

	w.Raw(instr)
}

func (w *AssemblyWriter) Print() {
	instr := `
	_print:
		STMFD R13!, {BP, LR}
		MOV BP, SP
		STMFD R13!, {R0, R1}

		LDR R0, [BP, #16]
		LDR R1, [BP, #12]

		BL printf

		MOV R0, #0

		BL fflush

		LDMFD R13!, {R0, R1}
		LDMFD R13!, {PC, BP}
	`

	w.Raw(instr)
}

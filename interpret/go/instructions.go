package main

const (
	val = iota
	rval
	lval
	rvar
	lvar
)

const (
	/* Standard instructions */
	nop = iota
	cmp
	cpy
	typ
	prt
	prv
	spr
	psh
	pop
	run
	get
	ret
	fun
	all
	del
	die
	err
	rea

	/* Jump */
	jmp
	jeq
	jne
	jgt
	jge
	jlt
	jle
	jer
	jnr

	/* Arithmetic */
	add
	sub
	mul
	div
	mod
	inc1
	inc2
	dec1
	dec2

	/* Bitwise */
	bor
	xor
	inv
	shl
	shr
	usr
)

type execFunc func()

type instr struct {
	args []int
	exec *execFunc
}

var instrNames = map[string][]int{
	"nop": []int{nop},
	"cmp": []int{cmp},
	"cpy": []int{cpy},
	"typ": []int{typ},
	"prt": []int{prt},
	"prv": []int{prv},
	"spr": []int{spr},
	"psh": []int{psh},
	"pop": []int{pop},
	"run": []int{run},
	"get": []int{get},
	"ret": []int{ret},
	"fun": []int{fun},
	"all": []int{all},
	"del": []int{del},
	"die": []int{die},
	"err": []int{err},
	"rea": []int{rea},

	"jmp": []int{jmp},
	"jeq": []int{jeq},
	"jne": []int{jne},
	"jgt": []int{jgt},
	"jge": []int{jge},
	"jlt": []int{jlt},
	"jle": []int{jle},
	"jer": []int{jer},
	"jnr": []int{jnr},

	"add": []int{add},
	"sub": []int{sub},
	"mul": []int{mul},
	"div": []int{div},
	"mod": []int{mod},
	"inc": []int{inc1, inc2},
	"dec": []int{dec1, dec2},

	"bor": []int{bor},
	"and": []int{and},
	"xor": []int{xor},
	"inv": []int{inv},
	"shl": []int{shl},
	"shr": []int{shr},
	"usr": []int{usr},
}

var instructions = []instr{
	/* nop */
	instr{
		args: []int{},
	},
	/* cmp */
	instr{
		args: []int{rval, rval},
	},
	/* cpy */
	instr{
		args: []int{lval, rval},
	},
	/* typ */
	instr{
		args: []int{lval, val},
	},
	/* prt */
	instr{
		args: []int{val},
	},
	/* prv */
	instr{
		args: []int{val},
	},
	/* spr */
	instr{
		args: []int{lval, rval},
	},
	/* psh */
	instr{
		args: []int{rval},
	},
	/* pop */
	instr{
		args: []int{lval},
	},
	/* run */
	instr{
		args: []int{rval, rvar},
	},
	/* get */
	instr{
		args: []int{lvar},
	},
	/* ret */
	instr{
		args: []int{rvar},
	},
	/* fun */
	instr{
		args: []int{lval, lvar},
	},
	/* all */
	instr{
		args: []int{lval, rval},
	},
	/* del */
	instr{
		args: []int{rval},
	},
	/* die */
	instr{
		args: []int{},
	},
	/* err */
	instr{
		args: []int{},
	},
	/* rea */
	instr{
		args: []int{lval, rval},
	},
	/* jmp */
	instr{
		args: []int{rval},
	},
	/* jeq */
	instr{
		args: []int{rval},
	},
	/* jne */
	instr{
		args: []int{rval},
	},
	/* jgt */
	instr{
		args: []int{rval},
	},
	/* jge */
	instr{
		args: []int{rval},
	},
	/* jlt */
	instr{
		args: []int{rval},
	},
	/* jle */
	instr{
		args: []int{rval},
	},
	/* jer */
	instr{
		args: []int{rval},
	},
	/* jnr */
	instr{
		args: []int{rval},
	},
	/* add */
	instr{
		args: []int{lval, rval, rval},
	},
	/* sub */
	instr{
		args: []int{lval, rval, rval},
	},
	/* mul */
	instr{
		args: []int{lval, rval, rval},
	},
	/* div */
	instr{
		args: []int{lval, rval, rval},
	},
	/* mod */
	instr{
		args: []int{lval, rval, rval},
	},
	/* inc1 */
	instr{
		args: []int{lval},
	},
	/* inc2 */
	instr{
		args: []int{lval, rval},
	},
	/* dec1 */
	instr{
		args: []int{lval},
	},
	/* dec2 */
	instr{
		args: []int{lval, rval},
	},
	/* bor */
	instr{
		args: []int{lval, rval, rval},
	},
	/* and */
	instr{
		args: []int{lval, rval, rval},
	},
	/* xor */
	instr{
		args: []int{lval, rval, rval},
	},
	/* inv */
	instr{
		args: []int{lval, rval, rval},
	},
	/* shl */
	instr{
		args: []int{lval, rval, rval},
	},
	/* shr */
	instr{
		args: []int{lval, rval, rval},
	},
	/* usr */
	instr{
		args: []int{lval, rval, rval},
	},
}

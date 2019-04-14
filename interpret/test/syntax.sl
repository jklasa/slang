; ===== IMPORTS =====
<<< "import.sl"
<</ "import2.sl" ; end of line comment

; ===== STRUCTURES =====
def struct (lat long)

def class (mem1 mem2) {

}

def class (mem2 mem2 mem3) {
    cpy idx 0

    def Function {
        ret
    }
}

def FunctionName (arg1 arg2 arg2) {
    cpy hello

    def localFunction (arg4):
        ret

    run localFunction

    ret (val1 val2 val3)
}

def Foo () {
    ret
}

def Printf () {
    ret
}

run Foo ()
run Printf

; ===== LABELS =====
#next:
#current
cpy idx 0 #end-of-line

; ===== VARIABLES =====
cpy one false
cpy two true
cpy three nil

cpy local >label
cpy Global 15

cpy str "this is a %s %v string \" still a string\n\\n"
cpy float -0.15214
cpy int 1452

cpy ptr1 &str
cpy ptr2 &[str + 1]
cpy drf1 *str
cpy drf2 *[str + &[1 + 1]]

def coords (lat long)
new location coords (124 153)
cpy temp location.lat

; ===== INSTRUCTIONS =====
nop
cmp idx 0
cpy idx 0
typ idx
prt idx
prv idx
spr dst idx
psh idx
pop idx
run @foo
run @foo (arg arg)
get arg2 arg
ret arg arg
ret
all lv rv
del lv
die
err
rea lv rv

jmp idx
jeq idx
jne idx
jgt idx
jge idx
jge idx
jlt idx
jle idx
jer idx
jnr idx

add lv rv1 rv2
sub lv rv1 rv2
mul l a b
div a b c
mod a b c
inc a
inc a b
dec a
dec a b
bor a b c
and a b c
xor a b c
inv a b c
shl a b c
shr a b c
usr a b c

def
new

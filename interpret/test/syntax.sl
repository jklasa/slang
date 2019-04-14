; ===== IMPORTS =====
<<< "import.sl"
<</ "import2.sl" ; end of line comment

; ===== FUNCTIONS =====
fun @function_name arg1 arg2 arg2:
    cpy hello
    ret val1 val2 val3

fun @foo:
    ret

fun @sl-printf:
    ret

run @foo
run @sl-printf

; ===== LABELS =====
#next:
#current
cpy idx 0 #end-of-line

; ===== VARIABLES =====
cpy one false
cpy two true
cpy three nil

cpy local 0
cpy $GLOBAL 15

cpy str "this is a %s %v string \" still a string\n\\n"
cpy float -0.15214
cpy int 1452

cpy ptr1 &str
cpy ptr2 &[str + 1]
cpy drf1 *str
cpy drf2 *[str + &[1 + 1]]

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
run @foo arg arg
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

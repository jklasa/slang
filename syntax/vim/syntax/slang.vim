" Vim syntax file
" Language: Slang
" Maintainer: Joel Klasa

if exists("b:current_syntax")
    finish
endif

syn keyword todo contained TODO FIXME XX NOTE
syn keyword instruction nop cmp cpy typ prv prt spr psh mod pop 
syn keyword instruction run get all del die err rea 
syn keyword instruction jmp jeq jne jgt jge jlt jle jer jnr 
syn keyword instruction add sub mul div inc dec 
syn keyword instruction bor and xor inv shl shr usr
syn keyword functionHeader fun ret

syn match local '[^ \.\[\]\t\n\]]\+'
syn match int '[-+]\=\d\+'
syn match float '[-+]\=\d*.\d\+'
syn match operator '[ \t][+-][ \t]'
syn match comment ';.*$' contains=TODO
syn match function '@[^ \t\n]*[:]*'
syn match libfunction '@sl-[^ \]\n\t]*[:]*'
syn match label '#[^ \[\] \t\n]*'
syn match label '>[^ \.\[\]\t\n]\+'
syn match import '<<<'
syn match import '<<\/'
syn match global '\$[^ \.\[\]\t\n]\+'

syn region string start="\"" skip="\\." end="\"" contains=TODO,escape
syn match escape '\\.'
syn match escape '%s'
syn match escape '%v'

let b:current_syntax = "slang"

hi def link functionHeader    Repeat
hi def link instruction       Type
hi def link string            String
hi def link escape            Delimiter
hi def link todo              Todo
hi def link comment           Comment
hi def link operator          Operator
hi def link function          Function
hi def link libfunction       Tag
hi def link label             Label
hi def link import            Include
hi def link int               Number
hi def link float             Float
hi def link global            Structure
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
syn keyword instruction def new

syn match int "\<-\=\d\+\%([Ee][-+]\=\d\+\)\=\>"
syn match float "\<-\=\d\+\.\d*\%([Ee][-+]\=\d\+\)\=\>"
syn match float "\<-\=\.\d\+\%([Ee][-+]\=\d\+\)\=\>"
syn match operator '\*'
syn match operator '&'
syn match comment ';.*$' contains=TODO
syn match label '#[^ \[\] \t\n]*'
syn match instruction '<<<'
syn match instruction '<<\/'
syn match global '\<[A-Z][A-Za-z0-9-_]*\>'

syn region string start="\"" skip="\\." end="\"" contains=TODO,escape
syn match escape '\\.'
syn match escape '%s'
syn match escape '%v'

let b:current_syntax = "slang"

hi def link instruction       Type
hi def link string            String
hi def link escape            Delimiter
hi def link todo              Todo
hi def link comment           Comment
hi def link operator          Operator
hi def link label             Label
hi def link int               Number
hi def link float             Float
hi def link global            Function

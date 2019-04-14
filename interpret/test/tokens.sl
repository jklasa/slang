fun @foo arg1 arg2 arg3:
    cpy idx 0.1124
    ret idx

; this is a comment
#label:
cpy idx 0 ; another comment
inc idx 2 #label
cpy string "hello world"
cpy other-string "|hello \" world \t\n|"

cpy int +1245
cpy float -1.21414

cpy expr0 ["hello world" - float]
cpy expr1 .[1.555555 + 12456]
cpy expr2 &[expr1 - expr2]
cpy expr3 [variable1 + [variable2 + variable3]]
# The Simple Language Specification

## Table of Contents
1. [Interpretation](#interpretation)
2. [Syntax](#syntax)
3. [Comments](#comments)
4. [Imports](#imports)
4. [Literals](#literals)
5. [Variables](#variables)
6. [Structures](#structures)
6. [Instructions](#instructions)
6. [Control Flow](#control-flow)
7. [Functions](#functions)
8. [Memory](#memory)

## Interpretation
The Simple language is an interpreted language, and code is read from input files into the interpreter's memory representation of the language. This representation is not fully specified and depends on the interpreter.

Running a Slang program is a two step process: first, the input program is read in its entirety into the bytecode in the loosest definition of the term, then this program as a whole is executed in the interpreter.

This detail is outside of the langauge itself, but the interpreter will use two types of errors: `InterpreterError` and `RuntimeError`. Interpreter errors will be errors encountered while reading and interpreting the supplied code. Usually the interpreter tries to find errors here before running the program so that time is not wasted waiting to reach errors that it should have known were going to happen anyway. Runtime errors are errors that the interpreter cannot find before execution. Instead, they are results from out of bounds accesses, invalid typing, invalid values, etc.

## Syntax
Slang is interpreted line by line, and each line is interpreted character by character. Whitespace is very important in this language and is used exclusively to differentiate keywords, instructions, variables, etc. in the language, so it is important that terms are separated appropriately. Indentation does not matter, so it is recommended to use it to format code for at least some readability. Newlines/lines also do not matter in the language (only whitespace does in general), so newlines can be used where desired to space out parts of code or declarations.

Functions, variables, labels, etc. are composed of only alphanumeric [A-Za-z0-9] characters, underscores _ and dashes -. Such named elements of the code must at least start with an alphabetic character. Certain other characters such as `[ ] " : ; ( ) { }` are reserved an cannot be used when not indicating variable expressions, string literals, and comments. For more detail about naming, see the section on [variables](#variables).

## Comments
Single-line comments exist in Slang and are delimited by semicolons anywhere in a line of code. Anything following the semicolon will be ignored by the interpreter.

Example
```
;--this is a comment--
cmp 1 0 ;----another comment here----
```

## Imports
Slang interprets and runs programs by combining multiple files into a single program. Files are stitched together by using code imports during the interpretation process (not during runtime).

There are two types of imports: guarded imports and unguarded imports. The syntax involves an import marker, determined by the import type, and a string specifying the relative path/location of the desired file to import. Naturally, neither type of import allows for recursive imports of files.

Imports can be placed anywhere in a file, but typical placement is at the beginning to keep things simple. An import will cause the interpretation of the current file to immediately pause and switch to interpretation of the new file. Think of imports as links to other files where, once reached, the contents of the imported file is copied into the current one before continuing reading the original file.


### Example
Below is a representation of how imports generally work. The import behavior is the focus here -- the code itself will make more sense after later explanation.

File 'main.sl'
```
cpy var 0
<<< "inc.sl"
prt var
```

File 'inc.sl'
```
add var var 5
```

Below is the resulting program that is executed. The variable `var` is set to the value `0`, after which it is incremented by 5. The final instruction prints this result with output `5`.

Program result:
```
cpy var 0
add var var 5
prt var
```

### Unguarded Imports

Syntax: `<<< "filename"`

Unguarded imports have default behavior where the import file is interpreted even if this file has already been imported. This is useful in situations where the import file is direct code instead of just functions for use.

### Guarded Imports

Syntax: `<</ "filename"`

Likely the more common type, guarded imports do not allow multiple imports of the same file. This type is typically used for files that only contain functions, e.g. library files. Similar to guards in C header files, using guarded imports allows the same import to be in multiple files without running into recursive import issues.

## Literals
There are three general types of literals in Slang: integers, floating-point values, and strings. Integer and float literals are exactly as one might expect. Strings are only delimited by double quotes, and any whitespace within the quotation marks will be preserved. With strings, there are two special characters: newline `\n` and tab `\t`. Both of these special characters will evaluate to newline and tab characters, respectively. Characters can be escaped via `\` to prevent the interpretation of special characters such as `\n`, `\t`, and `"`. INT and FLOAT literals do not actually have a specific location in memory. Instead, they are kept in the program representation in the interpreter, directly in the relevant instruction. STRING literals, on the other hand, are stored in stack frames, and the literal simply points to the start of the string in memory.

Examples:

| Literal | Type | Result |
| ------- | ---- | ------ |
| 123     | INT  | 123    |
| 123.0   | FLOAT | 123.0 |
| "Hello world" | STRING | Hello world |
| "Hello \\\" world" | STRING | Hello " world |

## Variables
Variables are all stored somewhere in the program's memory in the interpreter. Since there are no specific limitations on memory usage in a Slang program, there can be some creative usage of different variable types: all variables are just values and locations in memory, so they can generally all be used in the same manner.

Variables should be alphanumeric, with some possible symbols such as `-` and `_`. Other symbols should generally be avoided to prevent use of reserved markers and unknowingly changing a program's behavior. All variable names must begin with an alphabetic [A-Za-z] character.

There are differences for access of private or public variables at higher levels of scope, but all variables can be accessed in scope descendants.

### Private Variables
Private variables are names that the user can use to point to some location in memory. They can only be used after they have been declared/created via use as a left value in an [instruction](#instruction). Private variables have scope limited to the current [function](#function) block, and they cannot be used outside of the scope in which it was declared without error. The interpreter should ideally be able to catch illegal usage of private variables before execution.. Private variables are differentiated by a lowercase first alphabetic character in their names.

Private variables include standard variables that one might use in arithmetic, for example, as well as labels and structures. Labels define locations for jump instructions, so these must be limited to the local scope of the current structure. Structures that are set to be local via lowercase notation, are only visible to code within the current scope. Structures, as well as all other private variables, are not visible to outside code. They are, however, visible to scopes at deeper levels.

Private variables behave like local variables and private/unexported variables in usual programming language. The local variables part was previously discussed, so let's look at private/unexported behavior. There are things in this lanugage called structures, that basically behave like classes, functions, or structs. Private variables in structures are not exported and cannot be accessed outside of the structure. Alternatively, public variables can be accessed outside of the structure, and these variables will be discussed next.

Examples:

```
cpy variable 0
add another-variable variable 1

cpy Variable 0 ; this is NOT a private variable

def foo (arg1) { ; foo is a private function
    cpy loc 15

    def inner (arg2){
        cpy temp loc  ; this works since inner function is at a deeper scope than loc
        cpy below 0   ; this does not work because below is created later than inner
        ret
    }

    new below
    run foo
    ret
}

run inner ; this also does NOT work since inner is local to foo

def coordinates (Lat Long Distance) {
    def distance () {                      ; distance is a private function in the 
                                           ; class coordinates
        ret ([Lat + Long])
    }
}

new coordinates loc (15 30)
run coordinates.distance()      ; distance cannot be accessed because it is private

```

#### Public Variables
Public variables are differentiated from private variables via an uppercase first alphabetic character in their names. Like private variables, they must be created somewhere before usage, which the interpreter should detect before runtime. They behave like public/exported variables from other languages, meaning they are exported for access outside of the current structure. Labels cannot be public variables not matter the case notation of the label names. They are always local to only the current structure/scope/context.

Examples:
```
cpy GlobalVariable 1738

def foo (arg1) {

    #Label:    ; this label is still local to the current function
    cpy Global2 42

    def Inner(arg2) {
        ret
    }
    ret
}

run foo.Inner      ; this works since Inner is a public function
cpy temp foo.Global2   ; Global2 is a public variable, so it is accessible here
```

### Pointers
Aside from instruction keywords and numeric literals, nearly everything you see in the code has a location in memory. This would even include [structures](#structures) and [labels](#labels), but not variadics and expressions. In order to get the memory location of these variables, one can use a special marker `*` to dereference a pointer and access the value at the location pointed to by the pointer variable. The marker `&` can be used to obtain such a pointer from a variable itself. References and dereferences are compatible with any variable type, including structures and labels. Numeric literals, expressions, and variadics are not compatible with pointers and references/dereferences. While labels can be compatible with pointers, there is likely no good use case for this combination of features, so it might be best to avoid it.

Examples:
```
cpy variable 15
cpy pointer *variable    ; pointer points to the location of the INT literal 15
cpy deref &pointer       ; deref now has the value stored at the memory 
                         ; location pointed to by the pointer, e.g. 15

def foo (arg1) {
    ret
}

cpy funcptr &foo
run *foo          ; here we can run the function foo via a pointer to it
```

### Expressions
Variable expressions also exist in Slang, and they were created out of convenience for the user. Without them, tedious pointer arithmetic would frequently be needed to complete most tasks invovling memory, which happen to be most tasks. Variable expressions are essentially compound variables composed of other variables. They are delimited between square brackets `[]`, and the content inside the brackets is treated as if it were any other variable. Of course, this implies variable expressions can be used with references and dereferenes, and this is exaclty the use for which they were intended.

Expressions can fit a list of sums or differences of variable values. The format is `[var op var op var]`, where `var` is any variable (including another expressions), and `op` is either addition `+` or subtraction `-`. Expressions provide more power for relevant instructions by allowing additional arithmetic work to be performed within each instruction. It might be worth pointing out how certain addition or subtraction operations are more efficiently written and executed via variable epressions instead of the addition or subtraction instructions themselves.

Variable expressions are an older feature, so it might be useful to look into using structures before using explicit pointer work via variable expressions.

Examples
```
cpy variable [1 + 2 + 3]     ; variable = 6
cpy slow_var 1
add slow_var slow_var 2
add slow_var slow_var 3   ; much slower method for computing same answer of 6

; the below loop prints 10 digits from a string and uses a variable expression
; to replace pointer arithmetic that can be more cumbersome in select 
; situations
cpy str "0123456789"
cpy idx 0
cmp idx 10 #loop
jeq >done
prt *[str + idx]
inc idx
jmp >loop
#done

```

### Variadics
You might see variadic variables used in instructions that take a variable number of arguments or in function headers. They are generally some sequence of variables that are enclosed in parenthesis and are sent to the instruction as a group of variables. For further reading, head on over to [instructions](#instructions). These are not compatible with pointer references and dereferences and are only usable with instructions that accept them in that specific parameter location.

Examples:
```
def function (a b c) {
    ret (a b)
}

cpy idx 0
cpy str "hello world"
run function (idx str)
get (a b)
```

### Scopes
Scopes share characteristics with variadics. They are enclosures between braces that contain exclusively Slang code/instructions. They allow structures to contain code and other structures. They are not compatible with pointer references and dereferences and are only usable with instructions that accept them in that specific parameter location.

Since scopes are code enclosures, they have an implicit `ret` instruction with no arguments placed at the end of all code there. Then, the `ret` instruction is not mandatory in scopes, but they can be added explicitly for return values or returning somewhere other than the end of the scope.

Examples:
```
def function (a b c) {
    ret
}

def class (a b c) {
    def MemberFunction () {
        ret
    }
}
```

## Instructions
Slang programs will consist of mostly instructions. Instructions are defined by a select set of keywords that are all 3 characters in length for uniformity. They all take a specific number and type of arguments.

Traditionally, if there is more than one argument for an instruction, the leftmost value will be treated as a left value. Left values require the ability to be set. Such values include local variables, global variables, dereferences of pointers, labels, and functions. Basically most types except for literals and variable references. In the below instruction tables, left values are marked by `lv`.

Right values can be any variable type and generally have no restrictions. These values only require that they can be read, which is true for all variables and literals. In the below instruction tables, right values are marked by `rv`.

Most instructions require a specific number of arguments in a specific order, but some others allow for a variable number of arguments in groups. Options here are to supply no arguments in this position, or supply arguments normally, but contained within single left and right parentheses. These instructions will use the arguments available, provided they are valid in type and quantity.

The types of variables required within these groups depend on their use. In `def` structure definitions and `get` instructions, these variables must all be left values. Such variadic groups are labeled `(lvar)`. In `run` and `ret` instructions, the variables are all read and must all be right values. These variadic groups are labeled `(rvar)`.

Example of valid variadics in instructions:
```
run function                       ; empty variadic
run function ()                    ; empty variadic
run function (arg1)                ; variadic with one variable
run function (arg1 arg2 arg2)      ; variadic with three variables

def function {                     ; empty variadic

}
def function () {                  ; empty variadic

}
def function (arg1 arg2) {         ; varidadic with two variables

}
```

The last possible argument type is a scope. These were discussed in the variables section. They exclusively contain Slang instructions and code and are labeled `{code}` in the table.

Valid scopes in instructions:
```
def function () {           ; scope with instructions
    ret
}

def struct ()               ; empty scope

def class () {}             ; empty scope

def class () {              ; scope with instructions, including structures/scopes
    def class2 () {}
}
```

### Standard Instructions
| Instruction | Operand 1 | Operand 2 | Operand 3 | Behavior |
| ----------- | --------- | --------- | --------- | -------- |
| `nop`         | --        | --        | --        | Empty instruction, no behavior | 
| `cmp`         | `rv1`        | `rv2`        | --        | Compare `rv1 ?? rv2` and save the result flag. Used with jump instructions |
| `cpy`         | `lv`        | `rv`        | --        | Copy the value of `rv` into `lv`: `lv = rv` | 
| `typ`         | `lv`        | `r/lv`        | --        | Write type of value `r/lv` to `lv`: `0 -> INT, 1 -> FLOAT` | 
| `prv`         | `lv`        | --        | --        | Print the value itself at location `lv` |   
| `prt`         | `lv`        | --        | --        | Print value at location `lv` by treating it as a character | 
| `spr`         | `lv`        | `rv`        | --        | Write formatted string of the numeric value of `rv` into location set at `lv` | 
| `psh`         | `rv`        | --        | --        | Push rv onto the stack | 
| `pop`         | `lv`        | --        | --        | Pop value from top of stack in to `lv` | 
| `all`         | `lv`        | `rv`        | --        | Allocate a block of memory of size rv and store pointer to it in lv | 
| `del`         | `lv`        | --        | --        | Deallocate block of memory at location `lv` | 
| `die`         | --        | --        | --        | Immediately end the program | 
| `err`         | --        | --        | --        | Set error flag if flag is already false, else reset error flag | 
| `rea`         | `lv`        | `rv`        | --        | Read contents of filename, specified by string at `rv`, to the location at `lv` | 

### Structure Instructions
| Instruction | Operand 1 | Operand 2 | Operand 3 | Behavior |
| ----------- | --------- | --------- | --------- | -------- |
| `def`         | `lv`        | `(rvar)`        | `{code}`        | Define structure `lv` with members `(rvar)` and scope `{code}` |
| `run`         | `rv`        | `(rvar)`        | --        | Run function `rv`. Takes variable number of `r/lv` arguments in parentheses, all determined by the function being run, that are sent to the function as arguments |
| `get`         | `lv`        | `(lvar)`        | --        | Variable args. Set all left values equal to values popped from the stack, in reverse order. Intended to retrieve values returned from a function |
| `ret`         | `(rvar)`        | --        | --        | Return from current function. Variable supplied arguments specify values to be returned via the stack in opposite order. | 

### Jump Instructions
| Instruction | Operand 1 | Behavior |
| ----------- | --------- | -------- |
| `jmp`         | `rv`      | Unconditional jump to `rv` |
| `jeq`         | `rv`      | Jump to `rv` if `cmp` found `rv1 == rv2` | 
| `jne`         | `rv`      | Jump to `rv` if `cmp` found `rv1 != rv2` | 
| `jgt`         | `rv`      | Jump to `rv` if `cmp` found `rv1 > rv2` |   
| `jge`         | `rv`      | Jump to `rv` if `cmp` found `rv1 >= rv2` | 
| `jlt`         | `rv`      | Jump to `rv` if `cmp` found `rv1 < rv2` | 
| `jle`         | `rv`      | Jump to `rv` if `cmp` found `rv1 <= rv2` | 
| `jer`         | `rv`      | Jump to `rv` if error flag is set | 
| `jnr`         | `rv`      | Jump to `rv` if error flag is not set | 

| Instruction | Operand 1 | Operand 2 | Operand 3 | Behavior |
| ----------- | --------- | --------- | --------- | -------- |
| `jeq`       | `rv1`     | `rv2`     | `rv3`     | Jump to `rv3` if `rv1 == rv2` |
| `jne`       | `rv1`     | `rv2`     | `rv3`     | Jump to `rv3` if `rv1 != rv2` |
| `jgt`       | `rv1`     | `rv2`     | `rv3`     | Jump to `rv3` if `rv1 >  rv2` |
| `jge`       | `rv1`     | `rv2`     | `rv3`     | Jump to `rv3` if `rv1 >= rv2` |
| `jlt`       | `rv1`     | `rv2`     | `rv3`     | Jump to `rv3` if `rv1 <  rv2` |
| `jle`       | `rv1`     | `rv2`     | `rv3`     | Jump to `rv3` if `rv1 <= rv2` |

### Artihmetic Instructions
| Instruction | Operand 1 | Operand 2 | Operand 3 | Behavior |
| ----------- | --------- | --------- | --------- | -------- |
| `add`         | `lv`        | `rv1`        | `rv2`        | `lv = rv1 + rv2` |
| `sub`         | `lv`        | `rv1`        | `rv2`        | `lv = rv1 - rv2` |
| `mul`         | `lv`        | `rv1`        | `rv2`        | `lv = rv1 * rv2` |
| `div`         | `lv`        | `rv1`        | `rv2`        | `lv = rv1 / rv2` |
| `mod`         | `lv`        | `rv1`        | `rv2`        | `lv = rv1 % rv2` |
| `inc`         | `lv`        | --           | --           | `lv++` |
| `inc`         | `lv`        | `rv`         | --           | `lv += rv` |
| `dec`         | `lv`        | --           | --           | `lv--` |
| `dec`         | `lv`        | `rv`         | --           | `lv -= rv` |

### Bitwise Instructions
| Instruction | Operand 1 | Operand 2 | Operand 3 | Behavior |
| ----------- | --------- | --------- | --------- | -------- |
| `bor`         | `lv`        | `rv1`        | `rv2`        | bitwise or: <code>lv = rv1 &#124; rv2</code> |
| `and`         | `lv`        | `rv1`        | `rv2`        | bitwise and: `lv = rv1 & rv2` |
| `xor`         | `lv`        | `rv1`        | `rv2`        | bitwise xor: `lv = rv1 ^ rv2` |
| `inv`         | `lv`        | `rv`        | --        | negation: `lv = ~rv` | 
| `shl`         | `lv`        | `rv1`        | `rv2`        | left shift: `lv = rv1 << rv2` |
| `shr`         | `lv`        | `rv1`        | `rv2`        | right shift: `lv = rv1 >> rv2` |
| `usr`         | `lv`        | `rv1`        | `rv2`        | unsigned shift right: `lv = rv1 >>> rv2` |

## Structures
In a later addition to the language, data structures were added to simplify, well, a lot of things. Structues in Slang are more powerful than structs though: they are basically classes, ... or functions, ... or structs. They can have functions defined within them, have member variables, subclasses/inheritance, arguments, etc. They follow variable naming convention (including private/public behavior). They are defined using the `def` instruction and are created using the `new` instruction. Members of variables can be accessed via the dot `.` operator.

Structures are created and accessed exclusively via their pointers, so make sure to only use pointers with the dot operator.

Examples:
```

; ===== STRUCTs =====
def coordinates (lat long)         ; define local coordinates struct
new location coordinates (15 30)   ; allocate and initialize new coordinates
                                   ; location is a pointer to the memory block

cpy latitude location.lat          ; latitude = location.lat = 15
cpy longitude location.long        ; longitude = location.long = 30

; ===== CLASSES =====
def Example (mem1 mem2 Exported) {
    ; Code just placed inside a class is like constructor code
    cpy mem1 0
    cpy mem2 1

    ; Here is a public function structure
    def Sum () {
        ret ([mem1 + mem2])
    }
}

new Example c (5 2)   ; create new instance c of class Example using default values
run c                 ; run init/constructor code lying somewhere in the class
run c.Sum             ; run the Sum function using c's member variables
get (result)          ; get the result of the Sum function (result = 1)
cpy temp c.Exported   ; temp = c.Exported, which is 0 by default

; ===== FUNCTIONS =====
def function_name (arg1 arg2 arg2 ...) {
    ...
    ... instructions here
    ...
    ret (val1 val2 val3) ...
}

; Minimal required syntax for a function
def foo {
    ret
}

; Inner functions are possible
def outer () {
    def inner (arg) {
        def toofar () {
            ret
        }
        ret
    }
    ret
}

def Foo (a b c) {
    def inner () {
        ret ([b + c])
    }

    run inner
    get (partial)
    ret ([a + partial])
}

run Foo 1 2 3
get (sum)
```

As previously stated, structures can behave like functions, structs, or classes. Functions are structures intended to be run, structs are structures containing members and no scope or code, and classes are structures that contain members and some scope (which can be more members like structures or variables). Structures are declared using the `def` keyword followed by the structure name and some combination of variadic and scope arguments. Structures can be defined inside other functions if desired.

Structures follow private/public variable behavior. Private functions are defined with a lowercase first letter, and they are only accessible within scope of the same level or higher level. Both private and public structures are only viewable by name by code below where they are declared. Public structures differ in that they are accessible as members of parent structures. Public variables within structures are also available as members of their structures, whereas private variables are not.

Structures, like any other variable, can be overloaded, where the last variable to appear in the code will always overload previous variables. Overloading here is a little different than you might see elsewhere: do not count on Slang to track types of variables, members, return types, etc. There cannot be multiple variables of the same name: overloading only occurs by name, even if the relevant variables or structures are completely different types.

Again, the whitespace is generally ignored, so indentation is encouraged, especially in structure blocks. Structures can optionally take arguments and return values via variadics. This functionality can be replicated without using function arguments or return values, as shown below, but this is generally has not been as useful since arguments and return values have been added. Structures often contain scopes of code. As mentioned in the scopes section, since scopes are code enclosures, they have an implicit `ret` instruction with no arguments placed at the end of all code there. Then, the `ret` instruction is not mandatory in scopes, but they can be added explicitly for return values or returning somewhere other than the end of the scope.

```
; Use run and get arguments for the function (uses stack frames)
def bar-1 (a b) {
    cpy c a
    cpy d b
    ret c d
}

cpy a 1
cpy b 2
run bar-1 (a b)
get a b

; Explicitly use the stack to send and retrieve variables
def bar-2 {
    pop b
    pop a
    cpy c a
    cpy d b
    psh d
    psh c
    ret
}

cpy a 1
cpy b 2
psh a
psh b
run bar
pop c
pop d
```

When calling a structure as a function via the `run` instruction or when creating a new structure via the `new` instruction, usually you might supply the correct number of arguments, determined by the header of the defined structure. However, Slang does not care how many arguments or members are supplied to the structure. If there are fewer arguments supplied to the structure, the provided arguments will be sent, and any remaining arguments not specified by the caller will be set to the default value of `0` if these arguments are not defined in the structure. If these variables are defined somewhere in the structure, then they retain their previously-defined value. This is useful in the cases where you might want to have some default or optional arguments at the end of a structure's parameter list. Additionally, more arguments can be supplied to the structure. In this case, the extra arguments will be placed onto the user stack in reverse order. To obtain these arguments, simply make calls to `pop` where the first argument exceeding the list will be popped first. This has use cases similar to variadic arguments in C. e.g. the Slang library function `Printf`.

```
def class (x y z Distance) {
    def Distance () {
        ret ([x + y + z])
    }
}

new class c (3 4)  ; create new class with x=3, y=4, z=0, and Distance=Distance

def foo (a b) {
    prv a
    prv b
    ret
}

run foo ; uses default value 0 for both a and b

def bar (a) {
    pop b
    pop c
    ret
}

run bar (a b c)
```

## Control flow

Execution flow of a Slang program is controlled via [structure](#structures) functions or labels. Labels are a special type of variable, and they are defined using marker `#`. A label variable will store the instruction number at which it is defined. Notice that this is not necessarily the same as the line number if there are functions or empty lines in the program. Label declarations with a semicolon `:` after them (e.g. `#label:`) stores the next available instruction number instead of the current one. Basically, instead of pointing at ourselves here, we are pointing at code below the label. Labels can be declared on most lines with one exception: the lines on which functions are declared.

The point of including labels is to add more versatility to the use of jump instructions. Use labels as the values/locations to which we jump in the process of changing control flow. Notably, labels can only be used in the same function in which they are declared, but they can be declared before they are used, i.e. labels are another exception to this declaration/usage rule. The interpreter will still be able to catch use of labels that are never declared, however.

Labels are declared using their `#` marker for separation from the rest of the code, but they are accessed/read without this marker. If the label is declared via `#label`, they can be accessed via `label` OR via `>label`. The latter option is added so help differentiate betwween labels and other local variables since they are similar, but have enough differences to warrant some separation.

Examples:
```
#start:
cpy str "0123456789"
cpy idx 0
cmp idx 10 #loop
jeq >done

prt .[str + idx]
inc idx
jmp >loop

def foo {
    jmp >start  ; this is an invalid label access due to incompatible scope
    ret
}

#done:
run foo
```

## Memory
Memory in a Slang program is contained within the interpreter. There are unique, defined sections of memory, the details of which are irrelevant to the language itself. Each Memory location can store either an INT value or FLOAT value (Slang is not designed to be memory efficient).

### Stack
The user has access to their own interpreter-controlled stack of a fixed size. This used to be used explicitly for function argument and return values, but has since been replaced by language-dependent functionality. To add a value to the stack, use the `psh` instruction, and to pop a value from the stack, use the `pop` instruction. Errors will be thrown if the stack is empty on a pop or full on a push.

### Heap
For more-permanent storage of larger amounts of data, the heap can be used. The heap is controlled by the interpreter and will dynamically expand as needed. Blocks of memory can be allocated via `all` and deallocated via `del`. Allocated blocks will last in memory until deallocation, so blocks can be used across functions, but memory leaks can and will occur if you're not careful. It is also advised to stay within the bounds of the allocated blocks since metadata for maintaining the free and allocated blocks is stored within the heap next to the allocated heap blocks.

Calls to `new` for creation of structure instances will make implicit calls to `all` for allocating the appropriate memory in which to store the structures. These structures still need to be deallocated via `del` calls.

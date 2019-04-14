# The Simple Language Specification

## Table of Contents
1. [Interpretation](#interpretation)
2. [Characters](#characters)
3. [Comments](#comments)
4. [Imports](#imports)
5. [Variables](#variables)
6. [Instructions](#instructions)
6. [Control Flow](#control-flow)
7. [Functions](#functions)
8. [Memory](#memory)

## Interpretation
The Simple language is an interpreted language, and code is read from input files into the interpreter's memory representation of the language. This representation is not fully specified and depends on the interpreter.

Running a Slang program is a two step process: first, the input program is read in its entirety, then this program as a whole is executed in the interpreter.

This detail is outside of the langauge itself, but the interpreter will use two types of errors: `InterpreterError` and `RuntimeError`. Interpreter errors will be errors encountered while reading and interpreting the supplied code. Usually the interpreter tries to find errors here before running the program so that time is not wasted waiting to reach errors that it should have known were going to happen anyway. Runtime errors are errors that the interpreter cannot find before execution. Instead, they are results from out of bounds accesses, invalid typing, invalid values, etc.

## Characters
Slang is interpreted line by line, and each line is interpreted character by character. Whitespace is generally ignored, except when differentiating keywords and variable names and when reading string literals. So, it is recommended to use indentation of whatever type in function blocks to format code for at least some readability.

Functions, variables, labels, etc. are composed of only alphanumeric [A-Za-z0-9] characters, underscores _ and dashes -. Such named elements of the code must contain at least one alphabetic character. Certain other characters such as `[ ] " ; : ( )` are reserved an cannot be used when not indicating variable expressions, string literals, and comments.

## Comments
Comments are delimited by semicolons anywhere in a line of code. Anything following the semicolon will be ignored by the interpreter.

Example
```
;--this is a comment--
cmp idx 0 ;----another comment here----
```

## Imports
Slang interprets and runs programs by combining multiple files into a single program. Files are stitched together by using code imports; in fact, the only way to combine multiple files in the same program is via imports. It is not possible to provide multiple files as input to the interpreter. Only one 'main' file must be provided to the interpreter as input.

There are two types of imports: guarded imports and unguarded imports. The syntax involves an import marker, determined by the import type, and a string specifying the location of the desired file to import. Naturally, neither type of import allows for recursive imports of files.

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

## Variables
Variables are generally all stored somewhere in the program's memory in the interpreter. There are several different types, and most are differentiated via special markers before the variable names. These markers tell the interpreter which name-memory lookup table to use during the interpretation process. Since there are no specific limitations on memory usage in a Slang program, there can be some creative usage of different variable types: all variables are just values and locations in memory, so they can generally all be used in the same manner.

Variables should be alphanumeric, with some possible symbols such as `-` and `_`. Other symbols should generally be avoided to prevent use of reserved markers and unknowingly changing a program's behavior.

### Literals
There are three general types of literals in Slang: integers, floating-point values, and strings. Integer and float literals are exactly as one might expect. Strings are only delimited by double quotes, and any whitespace within the quotation marks will be preserved. With strings, there are two special characters: newline `\n` and tab `\t`. Both of these special characters will evaluate to newline and tab characters, respectively. Characters can be escaped via `//` to prevent the interpretation of special characters such as `\n`, `\t`, and `"`. INT and FLOAT literals do not actually have a specific location in memory. Instead, they are kept in the program representation in the interpreter. STRING literals, on the other hand, are stored in stack frames, and the literal simply points to the start of the string in memory.

Examples:

| Literal | Type | Result |
| ------- | ---- | ------ |
| 123     | INT  | 123    |
| 123.0   | FLOAT | 123.0 |
| "Hello world" | STRING | Hello world |
| "Hello \\\" world" | STRING | Hello " world |

### Local Variables
Local variables are names that the user can use to point to some location in memory. Local variables are the only variable type to not require a maker before the name, and generally speaking, consecutive characters forming one word will be treated as a name for the variable. Any variable can only be used after it has been declared/created via use as a right value in an [instruction](#instruction). Local variables have scope limited to the current [function](#function) block, and they cannot be used outside of the scope in which it was declared without error. The interpreter should ideally be able to catch illegal usage of local variables.

Examples:

```
cpy variable 0
add another-variable variable 1
```

#### Global Variables
Global variables are differentiated from local variables via the marker `$` at the beginning of the variable name. As long as a global variable has already been created via use as a right value, it can be referenced anywhere in the program, i.e. its scope is anywhere in the program.

Examples:
```
cpy 
```

### Pointers
Aside from instruction keywords, nearly everything you see in the code with the exception of numeric literals has a location in memory. This would even include [functions](#functions) and [labels](#labels). In order to get the memory location of these values/variables, one can use a special marker `*` to dereference a pointer and access the value at the location pointed to by the pointer variable. The marker `&` can be used to obtain such a pointer from a variable itself. References and dereferences are compatible with any variable type.

Examples:
```
cpy variable 15
cpy pointer *variable    ; pointer points to the location of the INT literal 15
cpy deref &pointer       ; deref now has the value stored at the memory 
                         ; location pointed to by the pointer, e.g. 15
```

### Variable Expressions
Variable expressions also exist in Slang, and they were created out of convenience for the user. Without them, tedious pointer arithmetic would frequently be needed to complete most tasks invovling memory, which happen to be most tasks. Variable expressions are essentially compound variables composed of other variables. They are delimited between square brackets `[]`, and the content inside the brackets is treated as if it were any other variable. Of course, this implies variable expressions can be used with references and dereferenes, and this is exaclty the use for which they were intended.

Expressions can fit a list of sums or differences of variable values. The format is `[var op var op var]`, where `var` is any variable (including another expressions), and `op` is either addition `+` or subtraction `-`. Expressions provide more power for relevant instructions by allowing additional arithmetic work to be performed within each instruction. It might be worth pointing out how certain addition or subtraction operations are more efficiently written and executed via variable epressions instead of the addition or subtraction instructions themselves.

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

### Labels and Functions as Variables
Like most elements of Slang, [labels](#control-flow) and [functions](#functions) are treated as variables as well. They can be read from as numeric (INT values), or written to as INT values. Attempts to write to these variable types with FLOAT values will sensibly result in error since these variables are as pointers to addresses. For more detail on these variable types (including range of scope), see their respective sections in the Slang spec.

### Variadic variables
You might see variadic variables used in instructions that take a variable number of arguments or in function headers. They are generally some sequence of variables that are enclosed in parenthesis and are sent to the instruction as a group of variables. For further reading, head on over to [instructions](#instructions).

Examples:
```
fun @function (a b c):
    ret (a b)

cpy idx 0
cpy str "hello world"
run @function (idx str)
get (a b)
```

## Instructions
Slang programs will consist of mostly instructions. Instructions, fitting one per line, are defined by a select set of keywords that are all 3 characters in length for uniformity. Most take a number of arguments that usually follow a general structure.

Traditionally, if there is more than one argument for an instruction, the leftmost value will be treated as a left value. Left values require the ability to be set. Such values include local variables, global variables, dereferences of pointers (pointers that are local, global, or variable expressions themselves), labels, and functions. Basically most types except for literals and variable references. In the below instruction tables, left values are marked by `lv`.

Right values can be any variable type and generally have no restrictions. These values only require that they can be read, which is true for all variables and literals. In the below instruction tables, right values are marked by `rv`.

Most instructions require a specific number of arguments in a specific order, but some others allow for a variable number of arguments in groups (indicated by `(r/lv ...)`). Options here are to supply no arguments in this position, or supply arguments normally, but contained within single parenthesis. These instructions will use the arguments available, provided they are valid in type and quantity.

The types of variables required within these groups depend on their use. In function headers and `get` instructions, these variables must all be left values. In `run` and `ret` instructions, the variables are all read and must all be right values.

Example:
```
run @function
run @function (arg1)
run @function (arg1 arg2 arg2)
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
| `run`         | `rv`        | (`r/lv` ...)        | --        | Run function `rv`. Takes variable number of `r/lv` arguments in parentheses, all determined by the function being run, that are sent to the function as arguments |
| `get`         | `lv`        | (`r/lv` ...)        | --        | Variable args. Set all left values equal to values popped from the stack, in reverse order. Intended to retrieve values returned from a function |
| `ret`         | (`r/lv` ...)        | --        | --        | Return from current function. Variable supplied arguments specify values to be returned via the stack in opposite order. | 
| `all`         | `lv`        | `rv`        | --        | Allocate a block of memory of size rv and store pointer to it in lv | 
| `del`         | `lv`        | --        | --        | Deallocate block of memory at location `lv` | 
| `die`         | --        | --        | --        | Immediately end the program | 
| `err`         | --        | --        | --        | Set error flag if flag is already false, else reset error flag | 
| `rea`         | `lv`        | `rv`        | --        | Read contents of filename, specified by string at `rv`, to the location at `lv` | 

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

## Functions

Functions are a big part of making most languages much more user friendly, so naturally, they were included in Slang. Functions are declared using the `fun` keyword followed by the function name and a set of arguments and a semicolon.

Functions cannot be defined inside other functions, but functions can access any other functions in the code, even before they are declared.

Functions can be overloaded, where the last function to appear in the code will always overload previous functions. Overloading here is a little different than you might see elsewhere: Slang does not track types of variables or function arguments or return types. Therefore, overloading only occurs by name, even if argument or return types or counts differ between functions. 

Syntax:
```
; All possible syntax for a function block
fun @function_name (arg1 arg2 arg2 ...):
    ...
    ... instructions here
    ...
    ret (val1 val2 val3) ...

; Minimal required syntax for a function
fun @foo:
    ret
```

Again, the whitespace is generally ignored, so indentation is encouraged, especially in function blocks like the one above. Functions can optionally take arguments and return values. This functionality can be replicated without using function arguments or return values, as shown below, but this is generally has not been as useful since arguments and return values have been added. Functions can only have 1 return instruction placed at its end since the return instructions themselves tell the interpreter when to change contexts out from the function.

```
; Use run and get arguments for the function (uses stack frames)
fun @bar-1 (a b):
    cpy c a
    cpy d b
    ret c d

cpy a 1
cpy b 2
run @bar-1 (a b)
get a b

; Explicitly use the stack to send and retrieve variables
fun @bar-2:
    pop b
    pop a
    cpy c a
    cpy d b
    psh d
    psh c
    ret

cpy a 1
cpy b 2
psh a
psh b
run @bar
pop c
pop d
```

When calling a function via the `run` instruction, usually you might supply the correct number of arguments, determined by the header of the called function. However, Slang does not care how many arguments are supplied to the function. If there are fewer arguments supplied to the called function, the provided arguments will be sent, and any remaining arguments not specified by the caller will be set to the default value of `0`. This is useful in the cases where you might want to have some default or optional arguments at the end of a function's parameter list. Additionally, more arguments can be supplied to the called function. In this case, the extra arguments will be placed onto the user stack in reverse order. To obtain these arguments, simply make calls to `pop` where the first argument exceeding the list will be popped first. This has use cases similar to variadic arguments in C. e.g. the Slang library function `sl-printf`.

```
fun @foo (a b):
    prv a
    prv b
    ret

run @foo ; uses default value 0 for both a and b

fun @bar a:
    pop b
    pop c
    ret

run @bar (a b c)
```

## Control flow

Execution flow of a Slang program is controlled [functions](#functions) or labels. Labels are a special type of variable, and they are defined using marker `#`. A label variable will store the instruction number at which it is defined. Notice that this is not necessarily the same as the line number if there are functions or empty lines in the program. Label declarations with a semicolon `:` after them (e.g. `#label:`) stores the next available instruction number instead of the current one. Basically, instead of pointing at ourselves here, we are pointing at code below the label. Labels can be declared on most lines with one exception: the lines on which functions are declared.

The point of including labels is to add more versatility to the use of jump instructions. Use labels as the values/locations to which we jump in the process of changing control flow. Notably, labels can only be used in the same function in which they are declared, but they can be declared before they are used.

To use labels, i.e. read from them, use a different marker `>`. This allows the interpreter to differentiate between labels being used for jumps/reading and label declarations on the same line.

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


#done:
run @foo

fun @foo:
    jmp >start  ; this is an invalid label access due to incompatible scope
    ret
```

## Memory
Memory in a Slang program is contained within the interpreter. There are unique, defined sections of memory, detailed in the table below. Each Memory location can store either an INT value or FLOAT value (Slang is not designed to be memory efficient).

| Segment | Start Location | Size | Storage |
| ------- | -------------- | ---- | ----- |
| GLOBALS | 1              | 199  | Global variables, 0 is reserved for NULL pointer |
| STACK FRAMES | 200 | 5000 | Instruction pointer, local variables, arguments |
| STACK | 5200 | 300 | User-accessible stack for `psh`, `pop` instructions |
| HEAP | 5500 | 6000 | Expandable, bulk block memory space via allocation and deletion |

### Stack
The user has access to their own interpreter-controlled stack of a fixed size. This used to be used explicitly for function argument and return values, but has since been replaced by language-dependent functionality. To add a value to the stack, use the `psh` instruction, and to pop a value from the stack, use the `pop` instruction. Errors will be thrown if the stack is empty on a pop or full on a push.

### Heap
For more-permanent storage of larger amounts of data, the heap can be used. The heap is controlled by the interpreter and will dynamically expand as needed. Blocks of memory can be allocated via `all` and deallocated via `del`. Allocated blocks will last in memory until deallocation, so blocks can be used across functions, but memory leaks can and will occur if you're not careful. It is also advised to stay within the bounds of the allocated blocks since metadata for maintaining the free and allocated blocks is stored within the heap next to the allocated heap blocks.

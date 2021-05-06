# B++ Documentation
B++ is a programming language initially developed by the developers for The Brain of TWOW Central. Check out the source code at [their GitHub repository](https://github.com/AeroAstroid/TheBrainOfTWOWCentral)! 

## Table of contents
- [Introduction](#introduction)
- [Hello, World!](#hello-world!)
- [Variables](#variables)
- [Basic Functions](#basic-functions)
- [Comparison](#comparison)
- [GOTO Statements](#goto-statements)
- [Builtin Functions](#builtin-functions)

## Introduction
In B++, everything is a tag. A tag is made of square brackets, with a function call in them! You can also provide tags as input to another tag. Arguments to a tag are seperated by spaces. For example:
```bpp
[MATH 5 * 7]
[CONCAT "hello w" "orld"]
[IF [COMPARE 6 != 4] "6 is not 4" "6 is 4"]
```
You can also do comments using a "#". For example:
```bpp
# This is a comment.
```

## Hello, World!
In B++, the return value of a tag is automatically printed. That means that, to make a hello, world!, you just need to do:
```bpp
"Hello, World!"
```

## Variables
Variables are made using the DEFINE and VAR statements. To define a variable, use:
```bpp
[DEFINE helloworld "Hello, World!"]
```
You can also use DEFINE to change a variable. 

To get the value of a variable, use 
```bpp
[VAR helloworld]
```
We can make a hello world program using variables by doing:
```bpp
[DEFINE helloworld "Hello, World!"]
[VAR helloworld]
```

## Data Types
B++ is a type-safe language. There are 4 types in B++:
- Strings (words/letters)
- Integers (whole numbers)
- Floats (decimals)
- Arrays (lists)

### Strings
Strings can be defined like any other variable. You can get a letter of a string using the INDEX function. For example:
```bpp
[INDEX "Hi!" 0]
```
This gets the first letter of the string "Hi!", or "H". Note that the first letter has an index of 0.

### Floats and Integers
You can do math on floats an integers, using [the MATH function](#math). 
Integers are defined by doing:
```bpp
[DEFINE a 7]
```
Floats are defined by doing:
```bpp
[DEFINE b 0.21]
```

### Arrays
Arrays can have any type as values. You can even store an array in an array! Define arrays using the ARRAY function. For example, the following program makes an array with the values 1, 2, 3 and 4:
```bpp
[ARRAY 1 2 3 4]
```
You can get a value at an index using.
```bpp
[INDEX [ARRAY 1 2 3 4] 0]
```
This gets the first element in the array. Note that the first element has an index of 0.

## Basic Functions
There are a few basic functions which you will most likely ue a lot. They are explained below.

### Math
The first one is math. To do math, simply use the MATH tag with a value, operator, and another value. For example:
```bpp
[MATH 100 + 100]
```
Supported operators are:

| Operator | Math Function |
| --- | --- |
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `^` | Power |

### String Formatting
To format a string, you would use the CONCAT function. This accepts any number of strings and concatenates them. For example:
```bpp
[CONCAT "Hello" ", " "World" "!"]
```
Prints "Hello, World!".

## Comparison
To compare values, you use the COMPARE function. For example:
```bpp
[COMPARE 6 = 4]
``` 
In B++, there aren't booleans. COMPARE just returns 1 if true, and 0 if false. 

B++ Supports many comparison operators:
| Operator | What it Does |
| --- | --- |
| `=` | Equals |
| `!=` | Not Equal |
| `>` | Greater Than |
| `<` | Less Than |
| `>=` | Greater Than or Equal To |
| `<=` | Less Than or Equal To |

If statements are ternary. Simply just do:
```bpp
[IF [COMPARE 6 != 4] "6 is not 4" "6 is 4"]
```
To make an if statement. To have more than one instruction in an IF statement, check out [GOTOs](#goto-statements).

## GOTO Statements
GOTO Statements allow branches and loops.

### Basic GOTO
```bpp
[GOTO a]
"This line will be skipped :("

[SECTION a]
"Hi!"
```
When you run this, you will notice that it only prints "Hi!". This is because the GOTO statement moved to the line with the section named "a".

### Fancy If Statements
Using this, we can execute multiple lines of code in an IF statement.
For example:
```bpp
[IF [COMPARE 1 = 1] [GOTO tmp1] [GOTO tmp2]]

[SECTION true]
"Yay! 1 is equal to 1!"
[GOTO endif]

[SECTION false]
"1 isn't equal to one?"
[GOTO endif]

[SECTION endif]
```
Let's go through this program. When 1 is equal to 1, it goes to the section called "true". In there, it has a print statement. Then, it goes to the section called "endif". This allows it to skip over the else part.

### Loops
We can also make loops using this. For example, let's make a loop that will print all the numbers up to 10:
```bpp
[DEFINE i 1]

[SECTION loop]

[VAR i] # Print the number

[DEFINE i [MATH [VAR i] + 1]] # Increase the number by 1
[IF [COMPARE [VAR i] < 10] [GOTO loop] ""] # Loop back to the start
```
In this program, it goes to a previous section until a requirement is satisfied.

## Builtin Functions
B++ has many builtin functions, which are listed below.

| Function Signature | Description |
| --- | --- |
| `[CHOOSE val]` | Gets a random index of `val`, which can be an array or a string. |
| `[REPEAT val n]` | Repeats the contents of a string or array `val`, `n` times. |
| `[RANDINT lower upper]` | Gets a random integer within the range `lower`, `upper`. |
| `[RANDOM lower upper]` | Gets a random float in the range `lower`, `upper`. |
| `[FLOOR val]` | Gets the floor, or rounds down float `val`. |
| `[CEIL val]` | Gets the ceiling, or rounds up float `val`. |
| `[ROUND val]` | Rounds float `val` to the nearest integer, or whole number. |
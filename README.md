# TildePlusPlus
A small esolang built in Go, where each time you use a goto statement, the flow of the program inverts.

In this repository there are multiple example ~++ programs: hello world, forRange10 (counts down from 10), and 2 fibonacci sequence variants (with/without code commments)
To use ~++, build this Go project and run the exe, then follow the instructions to view documentation (also attached in this README) or run .tplpl files

~++ DOCUMENTATION:
NOTES:
- code comments begin with: --
- in ~++, you can invert the flow of a program, for non-in-line comments that precede a block of code where the flow is inverted, it is conventional to end the comment with: ^
- leading/trailing white space on lines is allowed, however white space in the middle of lines will cause syntax errors
- in the following notes, x and y are used to represent integer line numbers - this is because to reference a variable in ~++, you must reference their line number


----- VARIABLES -----
- note: a variable declaration can only consist of an '=' sign and a raw value - not a reference to another value
DECLARING A NUMERIC VARIABLE:
- note: all numbers in ~++ have the underlying type: float64, meaning they can be any positive or negative floating point value
=5
REFERENCING A NUMERIC VARIABLE:
&x
EDITING A NUMERIC VARIABLE:
- note: editing variables in ~++ involves var1, followed by an operation, followed by var2 - both values must be references to predefined values
&x+&y -- this line adds &x to &y. other mathematical functions include: (+, -, *, /)
DECLARING A STRING VARIABLE:
='foo'
REFERENCING A STRING VARIABLE:
$x
STRING CONCATENATION:
- note: strings only support one function: (+), used for string concatenation
$x+$y -- this line concatenates $x to $y


----- PRINTING VARIABLES -----
- note: the print method in ~++ only accepts a reference to a single, predefined variable
PRINTING A NUMERIC VARIABLE:
>&x
PRINTING A STRING VARIABLE:
>$y


----- INVERT-FLOW-GOTO -----
- note: in ~++, using a goto command also inverts the flow of the program:
- if the interpreter is travelling down the program line by line when it reaches a goto statement, it will start travelling up the program line by line and vice versa
GOTO:
- note: goto statements can only accept raw values, not a reference to a variable
~x -- this will goto line x and invert the flow of the program


----- IF STATEMENTS -----
- note: in ~++, conditional branch statements are composed of a +/- sign, a numeric variable reference, a ?, and an invert-flow-goto statement
BRANCH IF >= 0:
+&x?~y --this line will invert-flow-goto line y if &x is greater than or equal to zero
BRANCH IF < 0:
-&x?~y --this line will invert-flow-goto line y if &x is less than zero


----- PROGRAM TERMINATION -----
! --this terminates a program - it is necessary for every ~++ program to have a termination point

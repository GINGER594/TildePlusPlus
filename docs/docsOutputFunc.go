package docsPkg

import (
	"fmt"
)

func PrintDocs() {
	fmt.Print(
		"" +
			"\n~++ DOCUMENTATION:\n" +
			"NOTES:\n" +
			"- code comments begin with: --\n" +
			"- in ~++, you can invert the flow of a program, for non-in-line comments that precede a block of code where the flow is inverted, it is conventional to end the comment with: ^\n" +
			"- leading/trailing white space on lines is allowed, however white space in the middle of lines will cause syntax errors\n" +
			"- in the following notes, x and y are used to represent integer line numbers - this is because to reference a variable in ~++, you must reference their line number\n\n\n" +
			"----- VARIABLES -----\n" +
			"- note: a variable declaration can only consist of an '=' sign and a raw value - not a reference to another value\n" +
			"DECLARING A NUMERIC VARIABLE:\n" +
			"- note: all numbers in ~++ have the underlying type: float64, meaning they can be any positive or negative floating point value\n=5\n" +
			"REFERENCING A NUMERIC VARIABLE:\n&x\n" +
			"EDITING A NUMERIC VARIABLE:\n" +
			"- note: editing variables in ~++ involves var1, followed by an operation, followed by var2 - both values must be references to predefined values\n" +
			"&x+&y -- this line adds &x to &y. other mathematical functions include: (+, -, *, /)\n" +
			"DECLARING A STRING VARIABLE:\n='foo'\n" +
			"REFERENCING A STRING VARIABLE:\n$x\n" +
			"STRING CONCATENATION:\n" +
			"- note: strings only support one function: (+), used for string concatenation\n" +
			"$x+$y -- this line concatenates $x to $y\n\n\n" +
			"----- PRINTING VARIABLES -----\n" +
			"- note: the print method in ~++ only accepts a reference to a single, predefined variable\n" +
			"PRINTING A NUMERIC VARIABLE:\n>&x\n" +
			"PRINTING A STRING VARIABLE:\n>$y\n\n\n" +
			"----- INVERT-FLOW-GOTO -----\n" +
			"- note: in ~++, using a goto command also inverts the flow of the program:\n" +
			"- if the interpreter is travelling down the program line by line when it reaches a goto statement, it will start travelling up the program line by line and vice versa\n" +
			"GOTO:\n" +
			"- note: goto statements can only accept raw values, not a reference to a variable\n" +
			"~x -- this will goto line x and invert the flow of the program\n\n\n" +
			"----- IF STATEMENTS -----\n" +
			"- note: in ~++, conditional branch statements are composed of a +/- sign, a numeric variable reference, a ?, and an invert-flow-goto statement\n" +
			"BRANCH IF >= 0:\n" +
			"+&x?~y --this line will invert-flow-goto line y if &x is greater than or equal to zero\n" +
			"BRANCH IF < 0:\n" +
			"-&x?~y --this line will invert-flow-goto line y if &x is less than zero\n\n\n" +
			"----- PROGRAM TERMINATION -----\n" +
			"! --this terminates a program - it is necessary for every ~++ program to have a termination point\n",
	)
}

package TPLPLInterpreterPkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// a function not bound to any struct - takes in an error, a line number & a line, then formats it into one error
func formatError(err error, lineNum int, line string) error {
	return fmt.Errorf("%s: line %d: %s", err.Error(), lineNum+1, line)
}

// interpreter struct for ~++ files
type TPLPLInterpreter struct {
	i                    int
	indexDir             int
	numMemory            map[int]float64
	strMemory            map[int]string
	currentProgramLength int
}

// takes in a ~++ program & processes it (tplpli : ~++ interpreter)
func (tplpli TPLPLInterpreter) Interpret(program []string) error {
	tplpli.resetProgramMemory(program)

	for {
		// returning an error if the current interpreter index is negative or longer than the file
		if tplpli.i < 0 || tplpli.i >= tplpli.currentProgramLength {
			return formatError(errors.New("line index outside of the current file"), tplpli.i, "")
		}

		line := program[tplpli.i]

		if len(line) < 1 { // blank lines
			tplpli.incrementIndex()
			continue
		}

		var err error
		switch opcode := line[0]; opcode {
		case '=':
			err = tplpli.declareVar(line)
		case '&':
			err = tplpli.editNumVar(line)
		case '$':
			err = tplpli.editStrVar(line)
		case '>':
			err = tplpli.printVar(line)
		case '~':
			err = tplpli.gotoLineNum(line) // invert-flow-goto
		case '+', '-':
			err = tplpli.evaluateConditional(line, rune(opcode)) // operand>=0 / operand<0 -> invert-flow-goto
		case '!':
			return nil // program end
		default:
			err = errors.New("syntax error") // other char = invalid syntax
		}

		if err != nil {
			return formatError(err, tplpli.i, line)
		}
	}
}

func (tplpli *TPLPLInterpreter) resetProgramMemory(program []string) {
	tplpli.i = 0
	tplpli.indexDir = 1
	tplpli.numMemory = map[int]float64{}
	tplpli.strMemory = map[int]string{}
	tplpli.currentProgramLength = len(program)
}

// moves the interpreter up/down the program by 1 line
func (tplpli *TPLPLInterpreter) incrementIndex() {
	tplpli.i += tplpli.indexDir
}

// takes in "=n" and defines a variable
func (tplpli *TPLPLInterpreter) declareVar(line string) error {
	operand := line[1:]
	if num, err := strconv.ParseFloat(operand, 64); err == nil { // checking for a number
		tplpli.numMemory[tplpli.i+1] = num
		tplpli.incrementIndex()
		return nil
	} else if len(operand) > 0 && string(operand[0]) == "'" && string(operand[len(operand)-1]) == "'" { // checking for a string
		tplpli.strMemory[tplpli.i+1] = strings.ReplaceAll(operand, "'", "")
		tplpli.incrementIndex()
		return nil
	} else { // if neither number or string (or another error was found) - syntax error
		return errors.New("syntax error")
	}
}

// gets variable from &n
func (tplpli TPLPLInterpreter) getNumVar(variable string) (float64, error) {
	if numIndex, err := strconv.Atoi(variable[1:]); err != nil || variable[0] != '&' {
		return 0, errors.New("syntax error")
	} else if num, ok := tplpli.numMemory[numIndex]; !ok {
		return 0, errors.New("variable does not exist")
	} else {
		return num, nil
	}
}

// gets variable from $n
func (tplpli TPLPLInterpreter) getStrVar(variable string) (string, error) {
	if strIndex, err := strconv.Atoi(variable[1:]); err != nil || variable[0] != '$' {
		return "", errors.New("syntax error")
	} else if str, ok := tplpli.strMemory[strIndex]; !ok {
		return "", errors.New("variable does not exist")
	} else {
		return str, nil
	}
}

// takes in "&n+&n" and edits the left value according to the expression
func (tplpli *TPLPLInterpreter) editNumVar(expr string) error {
	// parsing the expression
	opcode := ""
	elems := []string{}
	switch {
	case strings.Contains(expr, "+"):
		opcode = "+"
		elems = strings.Split(expr, "+")
	case strings.Contains(expr, "-"):
		opcode = "-"
		elems = strings.Split(expr, "-")
	case strings.Contains(expr, "*"):
		opcode = "*"
		elems = strings.Split(expr, "*")
	case strings.Contains(expr, "/"):
		opcode = "/"
		elems = strings.Split(expr, "/")
	default:
		return errors.New("syntax error: no valid function: (+, -, *, /) found in numerical expression")
	}

	// parsing the elements in the expression
	num1, err := tplpli.getNumVar(elems[0])
	if err != nil {
		return err
	}
	num1Index, _ := strconv.Atoi(elems[0][1:]) // err can be undefined here as the potential syntax error is already handled by getNumVar
	num2, err := tplpli.getNumVar(elems[1])
	if err != nil {
		return err
	}

	// evaluating the expression
	switch opcode {
	case "+":
		tplpli.numMemory[num1Index] = num1 + num2
	case "-":
		tplpli.numMemory[num1Index] = num1 - num2
	case "*":
		tplpli.numMemory[num1Index] = num1 * num2
	case "/":
		tplpli.numMemory[num1Index] = num1 / num2
	}

	tplpli.incrementIndex()
	return nil
}

// takes in "$n+$n" and edits the left value according to the expression
func (tplpli *TPLPLInterpreter) editStrVar(expr string) error {
	// parsing the expression
	if !strings.Contains(expr, "+") {
		return errors.New("syntax error: no valid function: (+) found in string concatenation expression")
	}

	//parsing the elements in the expression
	elems := strings.Split(expr, "+")
	str1, err := tplpli.getStrVar(elems[0])
	if err != nil {
		return err
	}
	str1Index, _ := strconv.Atoi(elems[0][1:]) // err can be undefined here as the potential syntax error is already handled by getStrVar
	str2, err := tplpli.getStrVar(elems[1])
	if err != nil {
		return err
	}

	// evaluating the expression
	tplpli.strMemory[str1Index] = str1 + str2

	tplpli.incrementIndex()
	return nil
}

// takes in print statements: ">&n", ">$n" and prints the variable
func (tplpli *TPLPLInterpreter) printVar(line string) error {
	operand := line[1:]
	if len(operand) < 2 || (operand[0] != '&' && operand[0] != '$') {
		return errors.New("syntax error")
	}
	if operand[0] == '&' {
		num, err := tplpli.getNumVar(operand)
		if err != nil {
			return err
		}
		fmt.Println(num)
	} else {
		str, err := tplpli.getStrVar(operand)
		if err != nil {
			return err
		}
		fmt.Println(str)
	}
	tplpli.incrementIndex()
	return nil
}

// takes in "~n" and performs invert-flow-goto
func (tplpli *TPLPLInterpreter) gotoLineNum(expr string) error {
	// checking if the ~n is formatted correctly (necessary as the function can be called from evaluateConditional())
	if len(expr) == 0 || expr[0] != '~' {
		return errors.New("syntax error")
	}
	lineNum, err := strconv.Atoi(expr[1:])
	if err != nil {
		return errors.New("syntax error")
	}
	tplpli.i = lineNum - 1 // -1 due to 0-based indexing (program[0] is line 1)
	tplpli.indexDir *= -1
	return nil
}

// takes in conditional expressions: "+&n?~n", "-&n?~n" and performs invert-flow-goto if the condition is fulfilled
func (tplpli *TPLPLInterpreter) evaluateConditional(line string, conditionType rune) error {
	// parsing the conditional
	if len(strings.Split(line, "?")) != 2 {
		return errors.New("syntax error")
	}

	// parsing conditional
	conditional := strings.Split(line[1:], "?") // parsing conditional, splits into: "&n", "~n"
	variable, resultantExpr := conditional[0], conditional[1]
	num, err := tplpli.getNumVar(variable)
	if err != nil {
		return err
	}

	// evaluating conditional
	if (conditionType == '+' && num >= 0) || (conditionType == '-' && num < 0) {
		err := tplpli.gotoLineNum(resultantExpr)
		if err != nil {
			return err
		}
	} else {
		tplpli.incrementIndex() // moving on if false
	}
	return nil
}

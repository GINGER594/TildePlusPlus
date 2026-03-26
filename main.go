package main

import (
	"fmt"
	"os"
	"strings"
	docsPkg "tplplenv/docs"
	TPLPLInterpreterPkg "tplplenv/interpreter"
)

// runs the ~++ env for running .tplpl programs
func main() {
	fmt.Println("\n-- welcome to the ~++ environment --\nd : docs, filepath (no speech-marks) : run .tplpl file, ! : quit")
	runEnv := true
	for runEnv {
		fmt.Print("\n> ")
		inp := ""
		fmt.Scan(&inp)

		switch inp {
		case "d":
			docsPkg.PrintDocs()
		case "!":
			runEnv = false
		default:
			if program, err := readProgramFile(inp); err != nil {
				fmt.Printf("> error occured when opening file %s: %s\n", inp, err.Error())
			} else {
				programInterpreter := TPLPLInterpreterPkg.TPLPLInterpreter{}
				err := programInterpreter.Interpret(program)
				if err != nil {
					fmt.Printf("\n> runtime error: %s\n", err.Error())
				}
			}
		}
	}
}

func readProgramFile(path string) ([]string, error) {
	if !strings.Contains(path, ".tplpl") {
		path += ".tplpl" // adding file extension if necessary
	}

	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	data := make([]byte, 1000000) // max chars = 1,000,000
	count, err := file.Read(data)
	if err != nil {
		return []string{}, err
	}

	return parseRawBytes(data[:count]), nil
}

// takes raw bytes, returns slice containing each line with comments, carriage-returns & leading/trailing whitespace removed
func parseRawBytes(rawBytes []byte) []string {
	program := strings.Split(string(rawBytes), "\n")
	for i := range program {
		program[i] = strings.TrimSpace(strings.ReplaceAll(strings.Split(program[i], "--")[0], "\r", ""))
	}
	return program
}

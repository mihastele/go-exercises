package worker

import (
	"bufio"
	"os"
	"strings"
)

type Result struct {
	Line    string
	LineNum int
	Path    string
}

type Results struct {
	Inner []Result
}

func NewResult(line string, lineNum int, path string) Result {
	return Result{Line: line, LineNum: lineNum, Path: path}
}

func FindInFile(path string, pattern string) *Results {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	results := Results{make([]Result, 0)}

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), pattern) {
			results.Inner = append(results.Inner, NewResult(scanner.Text(), lineNum, path))
		}
		lineNum++
	}
	if len(results.Inner) > 0 {
		return &results
	} else {
		return nil
	}
}

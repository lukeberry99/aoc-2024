package files

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strings"
)

// Read a file and return a slice of slices of strings
// where each inner slice represents a column of data
func ReadColumns(name string, numColumns int) [][]string {
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}
	lines := readLines(name, callingFile)

	columns := make([][]string, numColumns)
	for i := range columns {
		columns[i] = []string{}
	}

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != numColumns {
			panic("unexpected number of columns in line: " + line)
		}

		for i, part := range parts {
			columns[i] = append(columns[i], part)
		}
	}

	return columns
}

// Read a file and return a slice of strings, one for each line
func ReadLines(name string) []string {
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}
	return readLines(name, callingFile)
}

// Read a file and return a string containing the entire file
func Read(name string) string {
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}
	b, err := os.ReadFile(path.Join(path.Dir(callingFile), name))
	if err != nil {
		panic(err)
	}
	return string(b)
}

// read a file and return a slice of strings, one for each line
// a gap of one or more blank lines is used to split the file into groups
func ReadParagraphs(name string) [][]string {
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}
	lines := readLines(name, callingFile)
	var groups [][]string

	curGroup := make([]string, 0)

	for _, line := range lines {
		if line == "" {
			groups = append(groups, curGroup)
			curGroup = make([]string, 0)
		} else {
			curGroup = append(curGroup, line)
		}
	}

	if len(curGroup) > 0 {
		groups = append(groups, curGroup)
	}
	return groups
}

func readLines(name string, callingFile string) []string {
	inputFile, err := os.Open(path.Join(path.Dir(callingFile), name))
	if err != nil {
		panic(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			panic(err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

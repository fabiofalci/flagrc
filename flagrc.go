package flagrc

import (
	"bufio"
	"os"
	"strings"
)

func ProcessFlagrc(filePath string) {
	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			argument := strings.Trim(scanner.Text(), " ")
			argumentName := getOnlyArgumentName(argument)
			if !containsArgument(argumentName) {
				appendArgument(argument)
			}
		}
	}
}

func getOnlyArgumentName(line string) string {
	if index := strings.Index(line, "="); index > 0 {
		return line[0:index]
	}
	return line
}

func containsArgument(argumentName string) bool {
	for i, value := range os.Args {
		if i > 0 {
			if strings.HasPrefix(value, argumentName) {
				return true
			}
		}
	}
	return false
}

func appendArgument(argument string) {
	os.Args = append(os.Args, argument)
}

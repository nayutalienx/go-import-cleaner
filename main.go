package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		help()
	}
	argument := os.Args[1]
	if argument == "help" || argument == "-help" || argument == "--help" {
		help()
	}
	validateArgument(argument)
	code := getCodeFromFile(argument)
	processedCode := processImports(code)
	saveResultToFile(processedCode, argument)
}

func validateArgument(argument string) {
	if !strings.Contains(argument, ".go") {
		exitWithMessage("ERROR! only .go files allowed. \ntype go-import-cleaner help")
	}
}

func processImports(code string) string {
	if !strings.Contains(code, "import") {
		exitWithMessage("file does not have imports")
	}

	if strings.Contains(code, "import \"") {
		exitWithMessage("skipping file with single import without aliases")
	}

	if strings.Contains(code, "import (") || strings.Contains(code, "import(") {
		return processBracedImports(code)
	} else {
		return processSingleAliaseImport(code)
	}
}

func processSingleAliaseImport(code string) string {
	scanner := bufio.NewScanner(strings.NewReader(code))
	for scanner.Scan() {
		codeLine := scanner.Text()
		if strings.Contains(codeLine, "import") {
			splitedLine := strings.Split(codeLine, " ")
			processedLine := splitedLine[0] + " " + splitedLine[2]
			return strings.Replace(code, codeLine, processedLine, 1)
		}
	}
	exitWithMessage("failed processing single aliase import")
	return ""
}

func processBracedImports(code string) string {
	beforeImportKeyword := strings.Split(code, "import")
	if len(beforeImportKeyword) == 1 {
		exitWithMessage("there is nothing to clean")
	}

	importsOpenBrace := strings.Split(beforeImportKeyword[1], "(")
	importsCloseBrace := strings.Split(importsOpenBrace[1], ")")

	allImports := importsCloseBrace[0]

	processedImports := ""

	scanner := bufio.NewScanner(strings.NewReader(allImports))
	for scanner.Scan() {
		importLine := scanner.Text()
		if len(importLine) == 0 {
			continue
		}
		importLineSplitted := strings.Split(importLine, " ")
		if len(importLineSplitted) == 2 {
			processedImports += "\n" + strings.Replace(importLine, strings.TrimSpace(importLineSplitted[0])+" ", "", 1)
		} else {
			processedImports += "\n" + importLineSplitted[0]
		}
	}

	processedImports = deduplicateStrings(processedImports)

	return strings.Replace(code, allImports, processedImports+"\n", 1)
}

type StringSet map[string]int

func (m StringSet) Put(s string) {
	m[s] = 1
}

func (m StringSet) ToString() string {
	result := ""
	for k, _ := range m {
		result += "\n" + k + "\n"
	}
	return result
}

func deduplicateStrings(lines string) string {
	linesArray := strings.Split(lines, "\n")
	set := StringSet{}
	for _, l := range linesArray {
		set.Put(l)
	}
	return set.ToString()
}

func help() {
	exitWithMessage("go-import-cleaner path-to-go-file \nexample: go-import-cleaner main.go")
}

func exitWithMessage(message string, a ...interface{}) {
	if len(a) == 0 {
		fmt.Println(message)
	} else {
		fmt.Println(message, a)
	}
	os.Exit(0)
}

func getCodeFromFile(argument string) string {
	file, err := ioutil.ReadFile(argument)
	if err != nil {
		exitWithMessage("Read file error: ", err)
	}

	return string(file)
}

func saveResultToFile(code string, argument string) {
	processedFile, err := os.Create(argument)
	if err != nil {
		panic(err)
	}
	_, err = processedFile.WriteString(code)
	if err != nil {
		exitWithMessage("error when write result to file: ", err)
	}
	err = processedFile.Close()
	if err != nil {
		exitWithMessage("error when close result file: ", err)
	}
}

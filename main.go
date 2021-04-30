package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		help()
	}

	if len(os.Args) > 2 {
		exitWithMessage("only 1 argument allowed")
	}

	argument := os.Args[1]

	if argument == "help" || argument == "-help" || argument == "--help" {
		help()
	}

	if !strings.Contains(argument, ".go") {
		exitWithMessage("ERROR! only .go files allowed. \ntype go-import-cleaner help")
	}

	file, err := ioutil.ReadFile(argument)
	if err != nil {
		exitWithMessage("Read file error: ", err)
	}

	code := string(file)

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
			processedImports += "\n        " + importLineSplitted[1]
		} else {
			processedImports += "\n" + importLineSplitted[0]
		}
	}

	processedCode := strings.Replace(code, allImports, processedImports+"\n", 1)

	processedFile, err := os.Create(argument)
	if err != nil {
		panic(err)
	}
	_, err = processedFile.WriteString(processedCode)
	if err != nil {
		exitWithMessage("error when write result to file: ", err)
	}
	err = processedFile.Close()
	if err != nil {
		exitWithMessage("error when close result file: ", err)
	}

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
	os.Exit(1)
}

package main

import (
	"edit/edit"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) != 3 {
		fmt.Println("Wrong use of the tool")
		return
	}

	inputName := os.Args[1]
	outputName := os.Args[2]
	input, err := ioutil.ReadFile(inputName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	text := string(input)
	text = TextEditing(text)
	err = ioutil.WriteFile(outputName, []byte(text), 755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func TextEditing(text string) string {
	result := edit.EditHexBin(text)
	result = edit.EditCase(text)
	//result = edit.EditPunctuation(text)
	//result = edit.EditQuote(text)
	//result = edit.EditAN(text)
	return result
}

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func getPart(part string) string {
	fileName := strings.ReplaceAll(part, "@@", "")
	file, _ := os.ReadFile("html/" + fileName + ".html")
	return string(file)
}

func Build() {
	index, err := os.ReadFile("html/index.html")
	if err != nil {
		log.Fatal(err)
	}

	indexString := string(index)
	reg := regexp.MustCompile(`@@\w+`)
	matches := reg.FindAllString(indexString, -1)

	output := indexString
	for _, part := range matches {
		compiledPart := getPart(part)
		output = strings.ReplaceAll(output, part, compiledPart)
	}

	// fmt.Print(output)
	fmt.Println("Files builded...")

	wErr := os.WriteFile("./out/index.html", []byte(output), 0755)
	if wErr != nil {
		log.Fatal(wErr)
	}
}
func main() {
	Build()
}

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

func main() {
	index, err := os.ReadFile("html/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(index))

	indexString := string(index)

	reg := regexp.MustCompile(`@@\w+`)

	matches := reg.FindAllString(indexString, -1)

	// fmt.Println(matches)

	// Build final
	output := indexString

	for _, part := range matches {
		compiledPart := getPart(part)
		output = strings.ReplaceAll(output, part, compiledPart)
	}

	fmt.Print(output)

	wErr := os.WriteFile("./out/index.html", []byte(output), 755)
	if wErr != nil {
		log.Fatal(wErr)
	}
}

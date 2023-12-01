package main

import (
	"embed"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//go:embed data.txt
var f embed.FS

func main() {
	lines, err := f.ReadFile("data.txt")
	if err != nil {
		log.Println("data machine broke D:!", err)
		os.Exit(1)
	}

	count := 0

	for _, line := range strings.Split(string(lines), "\n") {
		if line == "" {
			continue
		}

		count += resolveNumberFromLine(line)
	}

	log.Println(count)
}

func resolveNumberFromLine(line string) (out int) {
	digitsInLine := []string{}

	for _, r := range line {
		if unicode.IsNumber(r) {
			digitsInLine = append(digitsInLine, string(r))
		}
	}

	if len(digitsInLine) == 0 {
		// ???
		return
	}

	join := digitsInLine[0] + digitsInLine[len(digitsInLine)-1]
	out, _ = strconv.Atoi(join)
	return
}

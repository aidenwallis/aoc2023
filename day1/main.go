package main

import (
	"embed"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
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

var transformMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

var matcher = regexp.MustCompile(`(` + strings.Join(maps.Keys(transformMap), "|") + `)`)

func resolveNumberFromLine(line string) (out int) {
	digits := []string{}

	if v := matcher.FindString(line); v != "" {
		digits = append(digits, transformMap[v])
	}

	for i := 0; i < len(line); i++ {
		// fucking cursed, walk back
		if v := matcher.FindString(line[len(line)-i:]); v != "" {
			digits = append(digits, transformMap[v])
			break
		}
	}

	// please dont do this
	out, _ = strconv.Atoi(digits[0] + digits[len(digits)-1])
	return
}

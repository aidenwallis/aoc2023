package main

import (
	"embed"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

//go:embed data.txt sample.txt
var f embed.FS

func main() {
	lines, err := f.ReadFile("data.txt")
	if err != nil {
		log.Println("data machine broke D:!", err)
		os.Exit(1)
	}

	powers := []int{}
	total := 0

	for _, line := range strings.Split(string(lines), "\n") {
		game := parseLine(line)
		if game == nil {
			slog.Error("failed to parse line", "line", line)
			continue
		}

		power := getPower(game)
		powers = append(powers, power)
		total += power
	}

	slog.Info("validGameIDs", "powers", powers, "total", total)
}

type Game struct {
	ID      int
	Subsets []*GameSubset
}

type GameSubset struct {
	Red   int
	Green int
	Blue  int
}

func parseLine(in string) *Game {
	spl := strings.SplitN(in, ":", 2)
	if len(spl) != 2 {
		return nil
	}

	gameHeader := strings.SplitN(spl[0], " ", 2)
	if len(spl) != 2 {
		return nil
	}

	gameID, err := strconv.Atoi(gameHeader[1])
	if err != nil {
		return nil
	}

	var subsets []*GameSubset
	for _, part := range strings.Split(spl[1], ";") {
		out := parseSubset(part)
		if out == nil {
			slog.Error("subset parsing failed", "part", part)
			continue
		}

		subsets = append(subsets, out)
	}

	return &Game{
		ID:      gameID,
		Subsets: subsets,
	}
}

func parseSubset(in string) *GameSubset {
	out := &GameSubset{}

	for _, part := range strings.Split(in, ",") {
		spl := strings.SplitN(strings.TrimSpace(part), " ", 2)
		if len(spl) != 2 {
			return nil
		}

		value, err := strconv.Atoi(spl[0])
		if err != nil {
			return nil
		}

		switch spl[1] {
		case "red":
			out.Red += value

		case "green":
			out.Green += value

		case "blue":
			out.Blue += value
		}
	}

	return out
}

func validateGame(game *Game) bool {
	for _, subset := range game.Subsets {
		if !validateSubset(subset) {
			return false
		}
	}
	return true
}

func getPower(game *Game) int {
	out := findMinimumCubes(game)
	return out.Red * out.Blue * out.Green
}

const (
	desiredRed   = 12
	desiredGreen = 13
	desiredBlue  = 14
)

func findMinimumCubes(game *Game) *GameSubset {
	totalRed := 0
	totalGreen := 0
	totalBlue := 0

	for _, subset := range game.Subsets {
		if subset.Red > totalRed {
			totalRed = subset.Red
		}
		if subset.Green > totalGreen {
			totalGreen = subset.Green
		}
		if subset.Blue > totalBlue {
			totalBlue = subset.Blue
		}
	}

	return &GameSubset{
		Red:   totalRed,
		Blue:  totalBlue,
		Green: totalGreen,
	}
}

func validateSubset(subset *GameSubset) bool {
	if subset.Red > desiredRed {
		return false
	}
	if subset.Blue > desiredBlue {
		return false
	}
	if subset.Green > desiredGreen {
		return false
	}
	return true
}

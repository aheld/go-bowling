package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	frames []Frame
}

type Frame struct {
	roll1 int
	roll2 int
}

func (f *Frame) Score() int {
	return f.roll1 + f.roll2
}

func (f *Frame) IsSpare() bool {
	return f.roll1+f.roll2 == 10
}

func (g *Game) Score() int {
	score := 0
	for i := 0; i < 10; i++ {
		score += g.frames[i].Score()
		if g.frames[i].IsSpare() {
			score += g.frames[i+1].roll1
		}
	}
	return score
}

func (g *Game) addFrame(roll1, roll2 int) {
	g.frames = append(g.frames, Frame{roll1, roll2})
}

func calcScoreForRollInput(rollInput string) int {
	game := Game{}
	rolls := strings.Split(rollInput, ",")
	for i := 0; i < len(rolls); i += 2 {
		rollAsInt1, _ := strconv.Atoi(rolls[i])
		rollAsInt2 := 0
		if i+1 < len(rolls) {
			rollAsInt2, _ = strconv.Atoi(rolls[i+1])
		}
		game.addFrame(rollAsInt1, rollAsInt2)
	}
	fmt.Println(game.frames)
	return game.Score()
}

func main() {
	firstArgument := os.Args[1]
	score := calcScoreForRollInput(firstArgument)
	fmt.Println(score)
}

package main

import (
	"fmt"
	"os"
)

func calcScoreForRollInput(rolls string) int {
	//score := 0
	//for _, roll := range strings.Split(rolls, ",") {
	//	rollAsInt, _ := strconv.Atoi(roll)
	//	score += rollAsInt
	//}
	//return score
	return -1
}

func main() {
	firstArgument := os.Args[1]
	score := calcScoreForRollInput(firstArgument)
	fmt.Println(score)
}

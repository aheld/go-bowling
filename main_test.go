package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name  string
		input string
		score int
	}{
		{
			name:  "Gutter Game",
			input: "0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0",
			score: 0,
		},
		//{
		//	name:  "simple Game",
		//	input: "1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1",
		//	score: 20,
		//},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			score := calcScoreForRollInput(test.input)
			if score != test.score {
				t.Errorf("Test %s failed, expected %v, got %v", test.name, test.score, score)
			}
		})
	}
}

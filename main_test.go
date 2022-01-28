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
		{
			name:  "simple Game",
			input: "1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1",
			score: 20,
		},
		{
			name:  "spare Game",
			input: "5,5,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0",
			score: 20,
		},
		{
			name:  "Last roll is a spare",
			input: "0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,9,1,5",
			score: 15,
		},
		// {
		// 	name:  "Role a strike",
		// 	input: "10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0",
		// 	score: 10,
		// },
		// {
		// 	name:  "Role a strike with a spare at the end",
		// 	input: "10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,5,5",
		// 	score: 25,
		// },
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

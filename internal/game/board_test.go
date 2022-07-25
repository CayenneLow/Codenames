package game

import (
	"encoding/json"
	"testing"

	"github.com/CayenneLow/Codenames/internal/game/enum"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	// 2x2 grid
	board := Board{
		Cells: [][]Cell{
			{
				Cell{
					Word:    "1x1",
					Team:    enum.BLUE_TEAM,
					Guessed: false,
				},
				Cell{
					Word:    "1x2",
					Team:    enum.RED_TEAM,
					Guessed: false,
				},
			},
			{
				Cell{
					Word:    "2x1",
					Team:    enum.BLUE_TEAM,
					Guessed: false,
				},
				Cell{
					Word:    "2x2",
					Team:    enum.RED_TEAM,
					Guessed: false,
				},
			},
		},
	}
	expected := `{
		"cells": [
			[
				{
					"word": "1x1",
					"team": 1,
					"guessed": "false"
				},
				{
					"word": "1x2",
					"team": 0,
					"guessed": "false"
				}
			],
			[
				{
					"word": "2x1",
					"team": 1,
					"guessed": "false"
				},
				{
					"word": "2x2",
					"team": 0,
					"guessed": "false"
				}
			]
		]
	}`
	actual, err := board.Json()
	assert.NoError(t, err)
	var expectedJson Board
	var actualJson Board
	json.Unmarshal([]byte(expected), &expectedJson)
	json.Unmarshal(actual, &actualJson)
	assert.Equal(t, expectedJson, actualJson)
}

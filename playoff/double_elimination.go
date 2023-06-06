// Copyright 2022 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Defines the tournament structure for a double-elimination bracket culminating in a best-of-three final.

package playoff

import (
	"fmt"
	"github.com/Team254/cheesy-arena/model"
)

// Creates a double-elimination bracket and returns the root matchup comprising the tournament finals. Only supports
// having exactly eight alliances.
func newDoubleEliminationBracket(numAlliances int) (*Matchup, error) {
	if numAlliances != 8 {
		return nil, fmt.Errorf("double-elimination bracket must have exactly 8 alliances")
	}

	// Define Round 1 matches.
	m1 := Matchup{
		id:                 "M1",
		NumWinsToAdvance:   1,
		redAllianceSource:  allianceSelectionSource{1},
		blueAllianceSource: allianceSelectionSource{8},
		matchSpecs:         newDoubleEliminationMatch(1, "Round 1 Upper"),
	}
	m2 := Matchup{
		id:                 "M2",
		NumWinsToAdvance:   1,
		redAllianceSource:  allianceSelectionSource{4},
		blueAllianceSource: allianceSelectionSource{5},
		matchSpecs:         newDoubleEliminationMatch(2, "Round 1 Upper"),
	}
	m3 := Matchup{
		id:                 "M3",
		NumWinsToAdvance:   1,
		redAllianceSource:  allianceSelectionSource{2},
		blueAllianceSource: allianceSelectionSource{7},
		matchSpecs:         newDoubleEliminationMatch(3, "Round 1 Upper"),
	}
	m4 := Matchup{
		id:                 "M4",
		NumWinsToAdvance:   1,
		redAllianceSource:  allianceSelectionSource{3},
		blueAllianceSource: allianceSelectionSource{6},
		matchSpecs:         newDoubleEliminationMatch(4, "Round 1 Upper"),
	}

	// Define Round 2 matches.
	m5 := Matchup{
		id:                 "M5",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m1, useWinner: false},
		blueAllianceSource: matchupSource{matchup: &m2, useWinner: false},
		matchSpecs:         newDoubleEliminationMatch(5, "Round 2 Lower"),
	}
	m6 := Matchup{
		id:                 "M6",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m3, useWinner: false},
		blueAllianceSource: matchupSource{matchup: &m4, useWinner: false},
		matchSpecs:         newDoubleEliminationMatch(6, "Round 2 Lower"),
	}
	m7 := Matchup{
		id:                 "M7",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m1, useWinner: true},
		blueAllianceSource: matchupSource{matchup: &m2, useWinner: true},
		matchSpecs:         newDoubleEliminationMatch(7, "Round 2 Upper"),
	}
	m8 := Matchup{
		id:                 "M8",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m3, useWinner: true},
		blueAllianceSource: matchupSource{matchup: &m4, useWinner: true},
		matchSpecs:         newDoubleEliminationMatch(8, "Round 2 Upper"),
	}

	// Define Round 3 matches.
	m9 := Matchup{
		id:                 "M9",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m7, useWinner: false},
		blueAllianceSource: matchupSource{matchup: &m6, useWinner: true},
		matchSpecs:         newDoubleEliminationMatch(9, "Round 3 Lower"),
	}
	m10 := Matchup{
		id:                 "M10",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m8, useWinner: false},
		blueAllianceSource: matchupSource{matchup: &m5, useWinner: true},
		matchSpecs:         newDoubleEliminationMatch(10, "Round 3 Lower"),
	}

	// Define Round 4 matches.
	m11 := Matchup{
		id:                 "M11",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m7, useWinner: true},
		blueAllianceSource: matchupSource{matchup: &m8, useWinner: true},
		matchSpecs:         newDoubleEliminationMatch(11, "Round 4 Upper"),
	}
	m12 := Matchup{
		id:                 "M12",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m10, useWinner: true},
		blueAllianceSource: matchupSource{matchup: &m9, useWinner: true},
		matchSpecs:         newDoubleEliminationMatch(12, "Round 4 Lower"),
	}

	// Define Round 5 matches.
	m13 := Matchup{
		id:                 "M13",
		NumWinsToAdvance:   1,
		redAllianceSource:  matchupSource{matchup: &m11, useWinner: false},
		blueAllianceSource: matchupSource{matchup: &m12, useWinner: true},
		matchSpecs:         newDoubleEliminationMatch(13, "Round 5 Lower"),
	}

	// Define final matches.
	final := Matchup{
		id:                 "F",
		NumWinsToAdvance:   2,
		redAllianceSource:  matchupSource{matchup: &m11, useWinner: true},
		blueAllianceSource: matchupSource{matchup: &m13, useWinner: true},
		matchSpecs:         newFinalMatches(14),
	}

	return &final, nil
}

// Helper method to create the matches for a given pre-final double-elimination matchup.
func newDoubleEliminationMatch(number int, nameDetail string) []*matchSpec {
	return []*matchSpec{
		{
			longName:            fmt.Sprintf("Match %d", number),
			shortName:           fmt.Sprintf("M%d", number),
			nameDetail:          nameDetail,
			order:               number,
			useTiebreakCriteria: true,
			tbaMatchKey:         model.TbaMatchKey{"sf", number, 1},
		},
	}
}

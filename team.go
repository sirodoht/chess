package main

// Team is the white or black team / player of a chess game
type Team int

const (
	// WHITE is the white team / player on the chess board
	WHITE Team = iota
	// BLACK is the black team / player on the chess board
	BLACK
	// NEITHER is when the square has no piece on it
	NEITHER
)

// GetTeamName returns the name of the team
func GetTeamName(team Team, format Format) string {
	if format == VERBOSE {
		verboseNames := map[Team]string{
			WHITE: "white ○",
			BLACK: "black ●",
		}
		return verboseNames[team]
	} else if format == SYMBOL {
		symbolNames := map[Team]string{
			WHITE: "○",
			BLACK: "●",
		}
		return symbolNames[team]
	} else if format == UPPER {
		upperNames := map[Team]string{
			WHITE: "WHITE",
			BLACK: "BLACK",
		}
		return upperNames[team]
	} else {
		lowerNames := map[Team]string{
			WHITE: "white",
			BLACK: "black",
		}
		return lowerNames[team]
	}
}

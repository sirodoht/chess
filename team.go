package main

// Team is the white or black team / player of a chess game
type Team int

const (
	// WHITE is the white team / player on the chess board
	WHITE Team = iota
	// BLACK is the black team / player on the chess board
	BLACK
)

// TeamFormat is in what format the team name should be returned
type TeamFormat int

const (
	// VERBOSE is e.g. "white ○"
	VERBOSE TeamFormat = iota
	// SYMBOL is only the circle symbol, e.g. "○"
	SYMBOL
	// UPPER is as uppercase letters
	UPPER
)

// GetTeamName returns the name of the team
func GetTeamName(team Team, format TeamFormat) string {
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
	} else {
		upperNames := map[Team]string{
			WHITE: "WHITE",
			BLACK: "BLACK",
		}
		return upperNames[team]
	}
}

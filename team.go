package main

// Team is the white or black team / player of a chess game
type Team int

const (
	// WHITE is the white team / player on the chess board
	WHITE Team = iota
	// BLACK is the black team / player on the chess board
	BLACK
)

// GetTeamName returns the name of the team
// e.g. WHITE -> "white ○"
func GetTeamName(team Team) string {
	TeamNames := map[Team]string{
		WHITE: "white ○",
		BLACK: "black ●",
	}
	return TeamNames[team]
}

// GetTeamSymbol returns the symbol / circle of the team
// e.g. WHITE -> "○"
func GetTeamSymbol(team Team) string {
	TeamNames := map[Team]string{
		WHITE: "○",
		BLACK: "●",
	}
	return TeamNames[team]
}

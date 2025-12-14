package models

type MapVeto struct {
	TeamName   string
	MapName    string
	TeamChoice string
}
type MatchSummary struct {
	MapChoices []MapVeto
}

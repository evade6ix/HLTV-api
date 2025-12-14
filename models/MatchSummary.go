package models

type MapVeto struct {
	TeamName   string
	MapName    string
	TeamChoice string
}

type TeamResult struct {
	TeamName string
	WonMaps  int
}
type MapResult struct {
	MapName string
	Team1   TeamResult
	Team2   TeamResult
}

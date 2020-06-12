package models

type Pick struct {
	ID             uint
	Manager        string `json:"manager"`
	PlayerName     string `json:"player_name"`
	PlayerPosition string `json:"player_position"`
	PlayerTeam     string `json:"player_team"`
}

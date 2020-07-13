package picks

type PickAPI struct {
	PlayerName string
	PlayerPosition string
	PlayerTeam string
}


func New() *PickAPI {
	return &PickAPI{}
} 


func (p *PickAPI) MakePick(picks *PickAPI) error {
	return nil
}

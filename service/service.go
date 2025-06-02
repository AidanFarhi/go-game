package service

import (
	"gogame/engine"
	"strings"
	"text/template"
)

type GameService struct {
	game engine.Game
}

func NewGameService(numRows, numCols int) GameService {
	return GameService{
		game: engine.NewGame(numRows, numCols),
	}
}

func (gs GameService) GetGrid() string {
	t, _ := template.ParseFiles("web/templates/grid.html")
	builder := strings.Builder{}
	t.Execute(&builder, gs.game.State)
	return builder.String()
}

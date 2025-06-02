package engine

import "errors"

type Cell struct {
	Occupied bool
	Color string
	Coordinates [2]int
}

type Game struct {
	NumRows int
	NumCols int
	State [][]Cell
}

func NewGame(numRows, numCols int) Game {
	game := Game{}
	game.NumRows = numRows
	game.NumCols = numCols
	for i := 0; i < numRows; i++ {
		row := []Cell{}
		for j := 0; j < numCols; j++ {
			row = append(row, Cell{false, "", [2]int{i, j}})
		}
		game.State = append(game.State, row)
	}
	return game
}

func (g Game) OccupyCell(row, col int, color string) error {
	if row >= g.NumRows || col >= g.NumCols {
		return errors.New("row or col out of bounds")
	}
	g.State[row][col].Occupied = true
	g.State[row][col].Color = color
	return nil
}

func (g Game) VacateCell(row, col int, color string) error {
	if row >= g.NumRows || col >= g.NumCols {
		return errors.New("row or col out of bounds")
	}
	g.State[row][col].Occupied = false
	g.State[row][col].Color = ""
	return nil
}

func (g Game) IsCellOccupied(row, col int) (bool, error) {
	if row >= g.NumRows || col >= g.NumCols {
		return false, errors.New("row or col out of bounds")
	}
	return g.State[row][col].Occupied, nil
}

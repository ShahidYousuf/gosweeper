package primary

import "gosweeper/internal/core/domain"

type BodyRevealCell struct {
	Row uint `json:"row"`
	Col uint `json:"col"`
}

type ResponseRevealCell domain.Game

func BuildResponseRevealCell(game domain.Game) ResponseRevealCell {
	return ResponseRevealCell(game)
}

package ports

import "gosweeper/internal/core/domain"

type GameRepository interface {
	Get(id string) (domain.Game, error)
	Save(game domain.Game) error
}

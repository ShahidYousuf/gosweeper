package ports

import "gosweeper/internal/core/domain"

type GameService interface {
	Get(id string) (domain.Game, error)
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, row uint, col uint) (domain.Game, error)
}

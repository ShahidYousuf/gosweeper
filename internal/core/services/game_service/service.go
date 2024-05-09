package game_service

import (
	"errors"
	"gosweeper/internal/core/domain"
	"gosweeper/internal/core/ports"
	"gosweeper/pkg"
)

type service struct {
	gamesRepository ports.GameRepository
	uid             pkg.UIDGen
}

func New(gameRepository ports.GameRepository, uidGen pkg.UIDGen) *service {
	return &service{
		gamesRepository: gameRepository,
		uid:             uidGen,
	}
}

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, pkg.NotFound) {
			return domain.Game{}, err
		}

		return domain.Game{}, errors.New("get game from repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New("the number of bombs is too high")
	}

	game := domain.NewGame(srv.uid.New(), name, size, bombs)

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("create game into repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Reveal(id string, row uint, col uint) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, pkg.NotFound) {
			return domain.Game{}, errors.New("game not found")
		}

		return domain.Game{}, errors.New("get game from repository has failed")
	}

	if !game.Board.IsValidPosition(row, col) {
		return domain.Game{}, errors.New("invalid position")
	}

	if game.IsOver() {
		return domain.Game{}, errors.New("game is over")
	}

	if game.Board.Contains(row, col, domain.CELL_BOMB) {
		game.State = domain.GAME_STATE_LOST
	} else {
		game.Board.Set(row, col, domain.CELL_REVEALED)

		if !game.Board.HasEmptyCells() {
			game.State = domain.GAME_STATE_WON
		}
	}

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("update game into repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

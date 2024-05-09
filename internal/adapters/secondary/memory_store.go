package secondary

import (
	"encoding/json"
	"errors"
	"gosweeper/internal/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{make(map[string][]byte)}
}

func (store *memkvs) Get(id string) (domain.Game, error) {
	if val, ok := store.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(val, &game)
		if err != nil {
			return domain.Game{}, errors.New("failed to get value from kvs")
		}
		return game, nil
	}
	return domain.Game{}, errors.New("game not found in kvs")
}

func (store *memkvs) Save(game domain.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New("failed to marshal game")
	}
	store.kvs[game.ID] = bytes
	return nil
}

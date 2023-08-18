package repository

import (
	"github.com/game-connect/gc-server/domain/model"
)

type GameRepository interface {
	FindByGameKey(gameKey string) (entity *model.Game, err error)
	List() (entity *model.Games, err error)
}
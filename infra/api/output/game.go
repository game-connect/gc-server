package output

import (
	"github.com/game-connect/gc-server/domain/model"
)

type ListGame struct {
	List    []ListGameContent `json:"list"`
	Message string             `json:"message"`
}

type ListGameContent struct {
	GameKey   string `json:"game_key"`
	GenreKey  string `json:"genre_key"`
	GameTitle string `json:"game_title"`
}

func ToListGame(c *model.Games) *ListGame {
	if c == nil {
		return nil
	}

	var list []ListGameContent
	for _, genre := range *c {
		genreContent := ListGameContent{
			GameKey:     genre.GameKey,
			GenreKey:    genre.GenreKey,
			GameTitle:   genre.GameTitle,
		}
		list = append(list, genreContent)
	}

	return &ListGame{
		List:    list,
		Message: "game list created",
	}
}

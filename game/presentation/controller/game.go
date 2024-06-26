package controller

import (
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/game/presentation/output"
	"github.com/game-connect/gc-server/game/presentation/response"
	"github.com/game-connect/gc-server/game/presentation/parameter"
)

type GameController interface {
	ListGenre() echo.HandlerFunc
	ListGenreKeys() echo.HandlerFunc
	ListGame() echo.HandlerFunc
	ListGameByAdminUserKey() echo.HandlerFunc
	ListGameKeys() echo.HandlerFunc
	ListGenreAndGame() echo.HandlerFunc
	CreateGame() echo.HandlerFunc
	DeleteGame() echo.HandlerFunc
}

type gameController struct {
	gameService service.GameService
}

func NewGameController(gameService service.GameService) GameController {
    return &gameController{
        gameService: gameService,
    }
}

// List
// @Summary     ジャンル一覧取得
// @tags        Genre
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGenre}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_genre [get]
func (gameController *gameController) ListGenre() echo.HandlerFunc {
	return func(c echo.Context) error {
		genreResult, err := gameController.gameService.ListGenre()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_genre", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListGenre(genreResult)
		response := response.SuccessWith("list_genre", 200, out)

		return c.JSON(200, response)
	}
}

// List
// @Summary     ジャンル一覧取得
// @tags        Genre
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGenreKeys}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_genre_keys [get]
func (gameController *gameController) ListGenreKeys() echo.HandlerFunc {
	return func(c echo.Context) error {
		genreResult, err := gameController.gameService.ListGenre()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_genre_keys", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListGenre(genreResult)
		response := response.SuccessWith("list_genre_keys", 200, out)

		return c.JSON(200, response)
	}
}

// List
// @Summary     ゲーム一覧取得
// @tags        Game
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_game [get]
func (gameController *gameController) ListGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		gameResult, err := gameController.gameService.ListGame()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToListGame(gameResult)
		response := response.SuccessWith("list_game", 200, out)

		return c.JSON(200, response)
	}
}

// @Summary     ゲーム一覧取得
// @tags        Game
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGameByAdminUserKey}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_game [get]
func (gameController *gameController) ListGameByAdminUserKey() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminUserKey := c.Param("adminUserKey")

		gameResult, err := gameController.gameService.ListGameByAdminUserKey(adminUserKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToListGameByAdminUserKey(gameResult)
		response := response.SuccessWith("list_game", 200, out)

		return c.JSON(200, response)
	}
}

// List
// @Summary     ゲーム一覧取得
// @tags        Game
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGameKey}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_game_keys [get]
func (gameController *gameController) ListGameKeys() echo.HandlerFunc {
	return func(c echo.Context) error {
		gameResult, err := gameController.gameService.ListGame()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game_keys", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToListGame(gameResult)
		response := response.SuccessWith("list_game_keys", 200, out)

		return c.JSON(200, response)
	}
}

// List
// @Summary     ジャンル＆ゲーム一覧取得
// @tags        Genre
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListGenreAndGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /genre/list_genre [get]
func (gameController *gameController) ListGenreAndGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		genreResult, err := gameController.gameService.ListGenre()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_genre_and_game", 500, out)

			return c.JSON(500, response)
		}

		gameResult, err := gameController.gameService.ListGame()
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_game_and_game", 500, out)
			
			return c.JSON(500, response)
		}

		out := output.ToListGenreAndGame(genreResult, gameResult)
		response := response.SuccessWith("list_genre_and_game", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     連携ゲーム作成
// @tags        LinkGame
// @Accept      json
// @Produce     json
// @Param       body body parameter.CreateGame true "連携ゲーム作成"
// @Success     200  {object} response.Success{items=output.CreateGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /game/{admin_user_key}/create_game  [post]
func (gameController *gameController) CreateGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		adminUserKey := c.Param("adminUserKey")
		gameParam := &parameter.CreateGame{}
		c.Bind(gameParam)

		linkGameResult, err := gameController.gameService.CreateGame(adminUserKey, gameParam)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("create_game", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToCreateGame(linkGameResult)
		response := response.SuccessWith("create_game", 200, out)

		return c.JSON(200, response)
	}
}

// Create
// @Summary     連携ゲーム削除
// @tags        LinkGame
// @Accept      json
// @Produce     json
// @Param       body body parameter.DeleteGame true "連携ゲーム削除"
// @Success     200  {object} response.Success{items=output.DeleteGame}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /link_game/{admin_user_key}/delete_game  [post]
func (gameController *gameController) DeleteGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		gameKey := c.Param("gameKey")

		err := gameController.gameService.DeleteGame(gameKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("delete_game", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToDeleteGame()
		response := response.SuccessWith("delete_game", 200, out)

		return c.JSON(200, response)
	}
}

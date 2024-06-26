package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/game-connect/gc-server/docs/swagger/game"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/game-connect/gc-server/log"
	"github.com/game-connect/gc-server/game/di"
)

func Init() {
	// di: wire ./game/di/wire.go
	adminUserController := di.InitializeAdminUserController()
	userController := di.InitializeUserController()
	gameController := di.InitializeGameController()
	gameScoreController := di.InitializeGameScoreController()
	gameUserController := di.InitializeGameUserController()

	userMiddleware := di.InitializeUserMiddleware()
	adminUserMiddleware := di.InitializeAdminUserMiddleware()
	gameMiddleware := di.InitializeGameMiddleware()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// genre: ジャンル関連
	genre := e.Group("/genre")
	genre.GET("/list_genre", gameController.ListGenre()) // genre/list_genre
	genre.POST("/list_genre_keys", gameController.ListGenreKeys()) // genre/list_genre_keys
	genre.GET("/list_game", gameController.ListGame()) // genre/list_game
	genre.POST("/list_game_keys", gameController.ListGameKeys()) // genre/list_game_keys
	genre.GET("/list_genre_and_game", gameController.ListGenreAndGame()) // genre/list_game_and_game

	// admin: 認証API
	admin := e.Group("/admin")
	admin.POST("/register_admin_user", adminUserController.RegisterAdminUser()) // admin/register_admin_use
	admin.POST("/login_admin_user", adminUserController.LoginAdminUser()) // admin/register_admin_user

	// admin: 認証済ユーザーのみアクセス可能
	admin.Use(adminUserMiddleware.CheckToken)
	admin.GET("/check_admin_user/:adminUserKey", adminUserController.CheckAdminUser()) // admin/register_admin_check/:adminUserKey
	admin.PUT("/logout_admin_user/:adminUserKey", adminUserController.LogoutAdminUser()) // admin/register_admin_logout/:adminUserKey
	admin.DELETE("/delete_admin_user/:adminUserKey", adminUserController.DeleteAdminUser()) // admin/register_admin_delete/:adminUserKey

	// link_game: ゲーム
	linkGame := e.Group("/link_game")
	linkGame.Use(adminUserMiddleware.CheckToken)
	linkGame.POST("/:adminUserKey/create_game", gameController.CreateGame()) // link_game/:adminUserKey/create_game
	linkGame.DELETE("/:adminUserKey/delete_game/:gameKey", gameController.DeleteGame()) // link_game/:adminUserKey/delete_game
	linkGame.GET("/:adminUserKey/list_game", gameController.ListGameByAdminUserKey()) // link_game/:adminUserKey/list_game

	// user: 認証API
	user := e.Group("/user")
	user.Use(gameMiddleware.CheckApiKey)
	user.POST("/login_user", userController.LoginUser()) // user/login_user

	// game: ゲームAPI
	game := e.Group("/game")
	game.Use(userMiddleware.CheckToken)
	game.POST("/update_game_score", gameScoreController.UpdateGameScore()) // game//:userKey/list_game_user
	game.GET("/:userKey/list_game_score/:gameKey", gameScoreController.ListGameScore()) // game/:userKey/list_game_score/:gameKey
	
	game.GET("/:userKey/list_game_user", gameUserController.ListGameUser()) // game//:userKey/list_game_user
	
	e.Logger.Fatal(e.Start(":8000"))
}

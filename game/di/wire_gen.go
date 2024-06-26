// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-connect/gc-server/config/database"
	"github.com/game-connect/gc-server/game/presentation/controller"
	"github.com/game-connect/gc-server/game/presentation/middleware"
	"github.com/game-connect/gc-server/game/service"
	"github.com/game-connect/gc-server/infra/dao"
)

// Injectors from wire.go:

// admin user
func InitializeAdminUserController() controller.AdminUserController {
	db := database.NewDB()
	adminUserRepository := dao.NewAdminUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	adminUserService := service.NewAdminUserService(adminUserRepository, transactionRepository)
	adminUserController := controller.NewAdminUserController(adminUserService)
	return adminUserController
}

// user
func InitializeUserController() controller.UserController {
	db := database.NewDB()
	gameUserRepository := dao.NewGameUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	userService := service.NewUserService(gameUserRepository, transactionRepository)
	userController := controller.NewUserController(userService)
	return userController
}

// game
func InitializeGameController() controller.GameController {
	db := database.NewDB()
	gameRepository := dao.NewGameDao(db)
	gameSettingRepository := dao.NewGameSettingDao(db)
	genreRepository := dao.NewGenreDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	gameService := service.NewGameService(gameRepository, gameSettingRepository, genreRepository, transactionRepository)
	gameController := controller.NewGameController(gameService)
	return gameController
}

// game score
func InitializeGameScoreController() controller.GameScoreController {
	db := database.NewDB()
	gameRepository := dao.NewGameDao(db)
	gameSettingRepository := dao.NewGameSettingDao(db)
	gameScoreRepository := dao.NewGameScoreDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	gameScoreService := service.NewGameScoreService(gameRepository, gameSettingRepository, gameScoreRepository, transactionRepository)
	gameScoreController := controller.NewGameScoreController(gameScoreService)
	return gameScoreController
}

// game user
func InitializeGameUserController() controller.GameUserController {
	db := database.NewDB()
	gameRepository := dao.NewGameDao(db)
	gameUserRepository := dao.NewGameUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	gameUserService := service.NewGameUserService(gameRepository, gameUserRepository, transactionRepository)
	gameUserController := controller.NewGameUserController(gameUserService)
	return gameUserController
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
	userMiddleware := middleware.NewUserMiddleware()
	return userMiddleware
}

// admin user
func InitializeAdminUserMiddleware() middleware.AdminUserMiddleware {
	db := database.NewDB()
	adminUserRepository := dao.NewAdminUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	adminUserService := service.NewAdminUserService(adminUserRepository, transactionRepository)
	adminUserMiddleware := middleware.NewAdminUserMiddleware(adminUserService)
	return adminUserMiddleware
}

// game
func InitializeGameMiddleware() middleware.GameMiddleware {
	db := database.NewDB()
	gameRepository := dao.NewGameDao(db)
	gameSettingRepository := dao.NewGameSettingDao(db)
	genreRepository := dao.NewGenreDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	gameService := service.NewGameService(gameRepository, gameSettingRepository, genreRepository, transactionRepository)
	gameMiddleware := middleware.NewGameMiddleware(gameService)
	return gameMiddleware
}

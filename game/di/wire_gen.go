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
	gameScoreRepository := dao.NewGameScoreDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	userService := service.NewUserService(gameUserRepository, gameScoreRepository, transactionRepository)
	userController := controller.NewUserController(userService)
	return userController
}

// link game
func InitializeLinkGameController() controller.LinkGameController {
	db := database.NewDB()
	linkGameRepository := dao.NewLinkGameDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	linkGameService := service.NewLinkGameService(linkGameRepository, transactionRepository)
	linkGameController := controller.NewLinkGameController(linkGameService)
	return linkGameController
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

// link game
func InitializeLinkGameMiddleware() middleware.LinkGameMiddleware {
	db := database.NewDB()
	linkGameRepository := dao.NewLinkGameDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	linkGameService := service.NewLinkGameService(linkGameRepository, transactionRepository)
	linkGameMiddleware := middleware.NewLinkGameMiddleware(linkGameService)
	return linkGameMiddleware
}

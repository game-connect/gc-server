// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/chat-connect/cc-server/api/presentation/controller"
	"github.com/chat-connect/cc-server/api/presentation/middleware"
	"github.com/chat-connect/cc-server/api/service"
	"github.com/chat-connect/cc-server/config/database"
	"github.com/chat-connect/cc-server/infra/dao"
)

// Injectors from wire.go:

// user
func InitializeUserController() controller.UserController {
	db := database.NewDB()
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	userService := service.NewUserService(userRepository, transactionRepository)
	userController := controller.NewUserController(userService)
	return userController
}

// room
func InitializeRoomController() controller.RoomController {
	db := database.NewDB()
	roomRepository := dao.NewRoomDao(db)
	roomUserRepository := dao.NewRoomUserDao(db)
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	roomService := service.NewRoomService(roomRepository, roomUserRepository, userRepository, transactionRepository)
	roomController := controller.NewRoomController(roomService)
	return roomController
}

// room_user
func InitializeRoomUserController() controller.RoomUserController {
	db := database.NewDB()
	roomRepository := dao.NewRoomDao(db)
	roomUserRepository := dao.NewRoomUserDao(db)
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	roomUserService := service.NewRoomUserService(roomRepository, roomUserRepository, userRepository, transactionRepository)
	roomUserController := controller.NewRoomUserController(roomUserService)
	return roomUserController
}

// channel
func InitializeChannelController() controller.ChannelController {
	db := database.NewDB()
	channelRepository := dao.NewChannelDao(db)
	chatRepository := dao.NewChatDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	channelService := service.NewChannelService(channelRepository, chatRepository, transactionRepository)
	channelController := controller.NewChannelController(channelService)
	return channelController
}

// chat
func InitializeChatController() controller.ChatController {
	db := database.NewDB()
	chatRepository := dao.NewChatDao(db)
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	chatService := service.NewChatService(chatRepository, userRepository, transactionRepository)
	chatController := controller.NewChatController(chatService)
	return chatController
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
	db := database.NewDB()
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	userService := service.NewUserService(userRepository, transactionRepository)
	userMiddleware := middleware.NewUserMiddleware(userService)
	return userMiddleware
}

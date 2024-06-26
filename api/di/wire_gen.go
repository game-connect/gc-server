// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-connect/gc-server/api/presentation/controller"
	"github.com/game-connect/gc-server/api/presentation/middleware"
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/config/database"
	"github.com/game-connect/gc-server/infra/dao"
)

// Injectors from wire.go:

// user
func InitializeUserController() controller.UserController {
	db := database.NewDB()
	userRepository := dao.NewUserDao(db)
	followRepository := dao.NewFollowDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	userService := service.NewUserService(userRepository, followRepository, transactionRepository)
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
	roomService := service.NewRoomService(roomRepository, roomUserRepository, userRepository, transactionRepository)
	followRepository := dao.NewFollowDao(db)
	roomUserService := service.NewRoomUserService(roomRepository, roomUserRepository, userRepository, followRepository, transactionRepository)
	roomUserController := controller.NewRoomUserController(roomService, roomUserService)
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

// follow
func InitializeFollowController() controller.FollowController {
	db := database.NewDB()
	followRepository := dao.NewFollowDao(db)
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	followService := service.NewFollowService(followRepository, userRepository, transactionRepository)
	followController := controller.NewFollowController(followService)
	return followController
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

// room_chat
func InitializeRoomChatController() controller.RoomChatController {
	db := database.NewDB()
	roomChatRepository := dao.NewRoomChatDao(db)
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	roomChatService := service.NewRoomChatService(roomChatRepository, userRepository, transactionRepository)
	roomChatController := controller.NewRoomChatController(roomChatService)
	return roomChatController
}

// open_chat
func InitializeOpenChatController() controller.OpenChatController {
	db := database.NewDB()
	openChatRepository := dao.NewOpenChatDao(db)
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	openChatService := service.NewOpenChatService(openChatRepository, userRepository, transactionRepository)
	openChatController := controller.NewOpenChatController(openChatService)
	return openChatController
}

// channel_chat
func InitializeChannelChatController() controller.ChannelChatController {
	db := database.NewDB()
	channelChatRepository := dao.NewChannelChatDao(db)
	userRepository := dao.NewUserDao(db)
	transactionRepository := dao.NewTransactionDao(db)
	channelChatService := service.NewChannelChatService(channelChatRepository, userRepository, transactionRepository)
	channelChatController := controller.NewChannelChatController(channelChatService)
	return channelChatController
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
	userMiddleware := middleware.NewUserMiddleware()
	return userMiddleware
}

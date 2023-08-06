package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/chat-connect/cc-server/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/chat-connect/cc-server/log"
	"github.com/chat-connect/cc-server/api/di"
)

func Init() {
	// di: wire ./api/di/wire.go
	userController := di.InitializeUserController()
	roomController := di.InitializeRoomController()
	roomUserController := di.InitializeRoomUserController()
	channelController := di.InitializeChannelController()
	chatController := di.InitializeChatController()

	userMiddleware := di.InitializeUserMiddleware()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth: 認証API
	auth := e.Group("/auth")
	auth.POST("/user_register", userController.UserRegister()) // auth/user_register
	auth.POST("/user_login", userController.UserLogin()) // auth/user_login

	// auth: 認証済ユーザーのみアクセス可能
	auth.Use(userMiddleware.UserMiddleware)
	auth.GET("/user_check/:userKey", userController.UserCheck()) // auth/user_check/:userKey
	auth.PUT("/user_logout/:userKey", userController.UserLogout()) // auth/user_logout/:userKey
	auth.DELETE("/user_delete/:userKey", userController.UserDelete()) // auth/user_delete/:userKey

	// room: 部屋関連
	room := e.Group("/room")
	room.Use(userMiddleware.UserMiddleware)
	room.GET("/:userKey/room_list", roomController.RoomList()) // room/:userKey/room_list
	room.POST("/:userKey/room_create", roomController.RoomCreate()) // room/:userKey/room_create
	room.DELETE("/:userKey/room_delete/:roomKey", roomController.RoomDelete()) // room/:userKey/room_delete/:roomKey

	room.POST("/:userKey/room_join/:roomKey", roomUserController.RoomJoin()) // room/:userKey/room_join/:roomKey
	room.DELETE("/:userKey/room_out/:roomKey", roomUserController.RoomOut()) // room/:userKey/room_out/:roomKey

	// channel: チャンネル関連
	channel := e.Group("/channel")
	channel.Use(userMiddleware.UserMiddleware)
	channel.GET("/:userKey/channel_list/:roomKey", channelController.ChannelList()) // channel/:userKey/channel_list/:roomKey
	channel.POST("/:userKey/channel_create/:roomKey", channelController.ChannelCreate()) // channel/:userKey/channel_create/:roomKey
	channel.DELETE("/:userKey/channel_delete/:channelKey", channelController.ChannelDelete()) // channel/:userKey/channel_delete/:channelKey

	// chat: チャット関連
	chat := e.Group("/chat")
	chat.Use(userMiddleware.UserMiddleware)
	chat.GET("/:userKey/chat_list/:channelKey", chatController.ChatList()) // chat/:userKey/chat_list/:channelKey
	chat.POST("/:userKey/chat_create/:channelKey", chatController.ChatCreate()) // chat/:userKey/chat_create/:channelKey

	e.Logger.Fatal(e.Start(":8000"))
}

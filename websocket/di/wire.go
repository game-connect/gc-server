// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-connect/gc-server/config/database"
    "github.com/game-connect/gc-server/infra/dao"
    "github.com/game-connect/gc-server/websocket/service"	
    "github.com/game-connect/gc-server/websocket/presentation/controller"
)

// chat
func InitializeChatController() controller.ChatController {
    wire.Build(
        database.NewDB,
        dao.NewChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewChatService,
        service.NewUserService,
        controller.NewChatController,
    )
    return nil
}

// channel_chat
func InitializeChannelChatController() controller.ChannelChatController {
    wire.Build(
        database.NewDB,
        dao.NewChannelChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewChannelChatService,
        service.NewUserService,
        controller.NewChannelChatController,
    )
    return nil
}

// open_chat
func InitializeOpenChatController() controller.OpenChatController {
    wire.Build(
        database.NewDB,
        dao.NewOpenChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewOpenChatService,
        service.NewUserService,
        controller.NewOpenChatController,
    )
    return nil
}

// room_chat
func InitializeRoomChatController() controller.RoomChatController {
    wire.Build(
        database.NewDB,
        dao.NewRoomChatDao,
        dao.NewUserDao,
        dao.NewTransactionDao,
        service.NewRoomChatService,
        service.NewUserService,
        controller.NewRoomChatController,
    )
    return nil
}

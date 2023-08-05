package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/model"
)

type ChannelRepository interface {
	ListByRoomKey(roomKey string) (entity *model.Channels, err error)
	Insert(channelModel *model.Channel, tx *gorm.DB) (entity *model.Channel, err error)
}
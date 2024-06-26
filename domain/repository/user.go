package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/model"
)

type UserRepository interface {
	FindByEmail(email string) (entity *model.User, err error)
	FindByName(name string) (entity *model.Users, err error)
	FindByUserKey(userKey string) (entity *model.User, err error)
	ListByUserKeys(userKeys []string) (entity *model.Users, err error)
	CountByStatus(status string) (count int64, err error)
	Insert(userModel *model.User, tx *gorm.DB) (entity *model.User, err error)
	Update(userModel *model.User, tx *gorm.DB) (entity *model.User, err error)
	DeleteByUserKey(userKey string, tx *gorm.DB) (err error)
}

package service

import (
	"log"

	"github.com/chat-connect/cc-server/domain/model"
	"github.com/chat-connect/cc-server/domain/repository"
	"github.com/chat-connect/cc-server/api/presentation/parameter"
	"github.com/chat-connect/cc-server/config/key"
)

type RoomService interface {
	RoomCreate(roomParam *parameter.RoomCreate, userKey string) (*model.Room, error)
}

type roomService struct {
	roomRepository        repository.RoomRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewRoomService(
		roomRepository repository.RoomRepository,
		userRepository repository.UserRepository,
		transactionRepository repository.TransactionRepository,
	) RoomService {
	return &roomService{
		roomRepository:        roomRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

// RoomCreate ルームを作成する
func (roomService *roomService) RoomCreate(roomParam *parameter.RoomCreate, userKey string) (roomResult *model.Room, err error) {
	// transaction
	tx, err := roomService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := roomService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := roomService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	userResult, err := roomService.userRepository.FindByUserKey(userKey)
	if err != nil {
		return nil, err
	}

	roomKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	roomModel := &model.Room{}
	roomModel.RoomKey = roomKey
	roomModel.UserID = userResult.ID
	roomModel.Name = roomParam.Name
	roomModel.Explanation = roomParam.Explanation
	roomModel.ImagePath = ""
	roomModel.UserCount = 1
	roomModel.Status = "public"

	roomResult, err = roomService.roomRepository.Insert(roomModel, tx)
	if err != nil {
		return nil, err
	}

	return roomResult, nil
}

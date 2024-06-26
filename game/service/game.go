package service

import (
	"log"

	"github.com/game-connect/gc-server/domain/model"
	"github.com/game-connect/gc-server/domain/repository"
	"github.com/game-connect/gc-server/game/presentation/parameter"
	"github.com/game-connect/gc-server/config/key"
	"github.com/game-connect/gc-server/infra/api"
)

type GameService interface {
	FindByApiKey(apiKey string) (*model.Game, error)
	ListGenre() (genreResult *model.Genres, err error)
	ListGame() (gameResult *model.Games, err error)
	ListGameByAdminUserKey(adminUserKey string) (gameResult *model.Games, err error)
	CreateGame(adminUserKey string, gameParam *parameter.CreateGame) (*model.Game, error)
	DeleteGame(gameKey string) (error)
}

type gameService struct {
	gameRepository        repository.GameRepository
	gameSettingRepository repository.GameSettingRepository
	genreRepository       repository.GenreRepository
	transactionRepository repository.TransactionRepository
}

func NewGameService(
		gameRepository        repository.GameRepository,
		gameSettingRepository repository.GameSettingRepository,
		genreRepository       repository.GenreRepository,
		transactionRepository repository.TransactionRepository,
	) GameService {
	return &gameService{
		gameRepository:        gameRepository,
		gameSettingRepository: gameSettingRepository,
		genreRepository:       genreRepository,
		transactionRepository: transactionRepository,
	}
}

// FindByApiKey api_keyからゲームを検索する
func (gameService *gameService) FindByApiKey(apiKey string) (*model.Game, error) {
	gameResult, err := gameService.gameRepository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	return gameResult, nil
}

// ListGenre ジャンル一覧を取得する
func (gameService *gameService) ListGenre() (genreResult *model.Genres, err error) {
	genreResult, err = gameService.genreRepository.List()
	if err != nil {
		return nil, err
	}

	return genreResult, nil
}

// ListGame ゲーム一覧を取得する
func (gameService *gameService) ListGame() (genreResult *model.Games, err error) {
	genreResult, err = gameService.gameRepository.List()
	if err != nil {
		return nil, err
	}

	return genreResult, nil
}

// ListGameByAdminUserKey 管理者ユーザーキーからゲーム一覧を取得する
func (gameService *gameService) ListGameByAdminUserKey(adminUserKey string) (genreResult *model.Games, err error) {
	genreResult, err = gameService.gameRepository.ListByAdminUserKey(adminUserKey)
	if err != nil {
		return nil, err
	}

	return genreResult, nil
}

// CreateGame 連携ゲームを作成する
func (gameService *gameService) CreateGame(adminUserKey string, gameParam *parameter.CreateGame) (*model.Game, error) {
	// transaction
	tx, err := gameService.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := gameService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := gameService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	gameKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	apiKey, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	gameModel := &model.Game{}
	gameModel.GameKey = gameKey
	gameModel.AdminUserKey = adminUserKey
	gameModel.ApiKey = apiKey
	gameModel.GameTitle = gameParam.GameTitle
	gameModel.GameImagePath = "/link_game/" + gameKey + ".png"
	gameModel.GenreKey = gameParam.GenreKey
	gameResult, err := gameService.gameRepository.Insert(gameModel, tx)
	if err != nil {
		return nil, err
	}

	gameSettingModel := &model.GameSetting{}
	gameSettingModel.GameKey = gameKey
	gameSettingModel.AdminUserKey = adminUserKey
	gameSettingModel.GameScore = gameParam.Setting.GameScore
	gameSettingModel.GameComboScore = gameParam.Setting.GameComboScore
	gameSettingModel.GameRank = gameParam.Setting.GameRank
	gameSettingModel.GamePlayTime = gameParam.Setting.GamePlayTime
	gameSettingModel.GameScoreImagePath = gameParam.Setting.GameScoreImage
	_, err = gameService.gameSettingRepository.Insert(gameSettingModel, tx)
	if err != nil {
		return nil, err
	}

	err = api.UploadImage(*gameParam.GameImage, gameKey, "/link_game")
	if err != nil {
		return nil, err
	}

	return gameResult, nil
}

// DeleteGame ユーザーを削除する
func (gameService *gameService) DeleteGame(gameKey string) (error) {
	// transaction
	tx, err := gameService.transactionRepository.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := gameService.transactionRepository.Rollback(tx)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := gameService.transactionRepository.Commit(tx)
			if err != nil {
				log.Panicln(err)
			}
		}
	}()

	err = gameService.gameRepository.DeleteByGameKey(gameKey, tx)
	if err != nil {
		return err
	}

	return nil
}

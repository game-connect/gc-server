package dao_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/chat-connect/cc-server/infra/dao"
	"github.com/chat-connect/cc-server/domain/model"
)

func TestRoomUserRepository_Insert(t *testing.T) {
    testCases := []struct {
        name             string
        mockParam        *model.RoomUser
        mockRowsAffected int64
        mockLastInsertID int64
        mockError        error
        expectedRoomUser *model.RoomUser
        expectedError    error
    }{
        {
            name: "Successful",
            mockParam: &model.RoomUser{
                RoomUserKey: "test_key",
                RoomKey:     "test_key",
                UserKey:     "test_key",
                Host:        false,
				Status:      "online",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedRoomUser: &model.RoomUser{
                ID:          1,
                RoomUserKey: "test_key",
                RoomKey:     "test_key",
                UserKey:     "test_key",
                Host:        false,
				Status:      "online",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
            },
            expectedError: nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            db, mock, _ := sqlmock.New()
            defer db.Close()

            gormDB, _ := gorm.Open("mysql", db)
            repo := dao.NewRoomUserRepository(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("INSERT").
                WillReturnResult(sqlmock.NewResult(tc.mockLastInsertID, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            user, err := repo.Insert(tc.mockParam, gormDB)
            if tc.expectedError != nil {
                assert.EqualError(t, err, tc.expectedError.Error())
            } else {
                assert.NoError(t, err)
            }

            user.CreatedAt = time.Time{}
            user.UpdatedAt = time.Time{}

            assert.Equal(t, tc.expectedRoomUser, user)
        })
    }
}

func TestRoomUserRepository_DeleteByRoomKeyAndUserKey(t *testing.T) {
    testCases := []struct {
        name            string
        roomKey         string
        userKey         string
        mockRowsAffected int64
        mockError       error
        expectedError   error
    }{
        {
            name:            "Successful",
            roomKey:         "test_key",
            userKey:         "test_key",
            mockRowsAffected: 1,
            mockError:        nil,
            expectedError:    nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            db, mock, _ := sqlmock.New()
            defer db.Close()

            gormDB, _ := gorm.Open("mysql", db)
            repo := dao.NewRoomUserRepository(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("DELETE").WithArgs(tc.userKey).
                WillReturnResult(sqlmock.NewResult(tc.mockRowsAffected, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            err := repo.DeleteByRoomKeyAndUserKey(tc.roomKey, tc.userKey, gormDB)
            assert.Equal(t, tc.expectedError, err)
        })
    }
}

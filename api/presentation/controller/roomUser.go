package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	
	"github.com/game-connect/gc-server/api/service"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
)

type RoomUserController interface {
	ListRoomUser() echo.HandlerFunc
	JoinRoom() echo.HandlerFunc
	OutRoom() echo.HandlerFunc
}

type roomUserController struct {
	roomService     service.RoomService
	roomUserService service.RoomUserService
}

func NewRoomUserController(
		roomService     service.RoomService,
		roomUserService service.RoomUserService,
	) RoomUserController {
    return &roomUserController{
		roomService:     roomService,
		roomUserService: roomUserService,
    }
}

// List
// @Summary     ルーム参加ユーザー一覧取得
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.ListRoomUser}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/list_room_user/{roomKey} [get]
func (roomUserController *roomUserController) ListRoomUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		userKey := c.Param("userKey")
		roomKey := c.Param("roomKey")

		roomUserResults, err := roomUserController.roomUserService.ListRoomUser(userKey, roomKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("list_room_user", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToListRoomUser(userKey, roomUserResults)
		response := response.SuccessWith("list_room_user", 200, out)

		return c.JSON(200, response)
	}
}

// Join
// @Summary     ルーム参加
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.JoinRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/room_join/{roomKey}  [post]
func (roomUserController *roomUserController) JoinRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")

		roomResult, err := roomUserController.roomService.FindByRoomKey(roomKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_join", 500, out)

			return c.JSON(500, response)
		}

		if roomResult == nil || roomResult.Status == "private" {
			out := output.NewError(fmt.Errorf("room does not exist"))
			response := response.ErrorWith("room_join", 400, out)

			return c.JSON(500, response)
		}

		roomUserResult, err := roomUserController.roomUserService.JoinRoom(roomKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_join", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToJoinRoom(roomUserResult)
		response := response.SuccessWith("room_join", 200, out)

		return c.JSON(200, response)
	}
}

// Out
// @Summary     ルーム退出
// @tags        Room
// @Accept      json
// @Produce     json
// @Success     200  {object} response.Success{items=output.OutRoom}
// @Failure     500  {object} response.Error{errors=output.Error}
// @Router      /room/{userKey}/room_out/{roomKey}  [delete]
func (roomUserController *roomUserController) OutRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		// parameters
		roomKey := c.Param("roomKey")
		userKey := c.Param("userKey")

		err := roomUserController.roomUserService.OutRoom(roomKey, userKey)
		if err != nil {
			out := output.NewError(err)
			response := response.ErrorWith("room_out", 500, out)

			return c.JSON(500, response)
		}

		out := output.ToOutRoom()
		response := response.SuccessWith("room_out", 200, out)

		return c.JSON(200, response)
	}
}

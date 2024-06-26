package output

import (
	"github.com/game-connect/gc-server/domain/dto"
	"github.com/game-connect/gc-server/domain/model"
)

type FollowingAndFollowers struct {
	UserKey        string `json:"user_key"`
	FollowingCount int64  `json:"following_count"`
	FollowersCount int64  `json:"followers_count"`
	Message        string `json:"message"`
}

func ToCountFollowingAndFollowers(userKey string, followingCount , followersCount int64) *FollowingAndFollowers {
	return &FollowingAndFollowers{
		UserKey: userKey,
		FollowingCount: followingCount,
		FollowersCount: followersCount,
 		Message: "count following adnd followers created",
	}
}

type ListFollows struct {
	UserKey string               `json:"user_key"`
	List    []ListFollowsContent `json:"list"`
	Message string               `json:"message"`
}

type ListFollowsContent struct {
	FollowKey        string      `json:"follow_key"`
	UserKey          string      `json:"user_key"`
	FollowingUserKey string      `json:"following_user_key"`
	Mutual           bool        `json:"mutual"`
	MutualFollowKey  string      `json:"mutual_follow_key"`
	Status           *UserStatus `json:"status"`
}

func ToListFollowing(userKey string, f *dto.FollowAndUsers) *ListFollows {
	if f == nil {
		return nil
	}

	var list []ListFollowsContent
	for _, followAndUser := range *f {
		followsContent := ListFollowsContent{
			FollowKey:        followAndUser.Follow.FollowKey,
			UserKey:          followAndUser.Follow.UserKey,
			FollowingUserKey: followAndUser.Follow.FollowingUserKey,
			Mutual:           followAndUser.Follow.Mutual,
			MutualFollowKey:  followAndUser.Follow.MutualFollowKey,
			Status:           ToUserStatus(&followAndUser.User),
		}
		list = append(list, followsContent)
	}

	return &ListFollows{
		UserKey: userKey,
		List:    list,
		Message: "list following created",
	}
}

func ToListFollowers(userKey string, f *dto.FollowAndUsers) *ListFollows {
	if f == nil {
		return nil
	}

	var list []ListFollowsContent
	for _, followAndUser := range *f {
		followsContent := ListFollowsContent{
			FollowKey:        followAndUser.Follow.FollowKey,
			UserKey:          followAndUser.Follow.UserKey,
			FollowingUserKey: followAndUser.Follow.FollowingUserKey,
			Mutual:           followAndUser.Follow.Mutual,
			MutualFollowKey:  followAndUser.Follow.MutualFollowKey,
			Status:           ToUserStatus(&followAndUser.User),
		}
		list = append(list, followsContent)
	}

	return &ListFollows{
		UserKey: userKey,
		List:    list,
		Message: "list followers created",
	}
}

type CreateFollow struct {
	FollowKey        string `json:"follow_key"`
	UserKey          string `json:"user_key"`
	FollowingUserKey string `json:"following_user_key"`
	Mutual           bool   `json:"mutual"`
	MutualFollowKey  string `json:"mutual_follow_key"`
	Message          string `json:"message"`
}

func ToCreateFollow(f *model.Follow) *CreateFollow {
	if f == nil {
		return nil
	}

	return &CreateFollow{
		FollowKey:        f.FollowKey,
		UserKey:          f.UserKey,
		FollowingUserKey: f.FollowingUserKey,
		Mutual:           f.Mutual,
		MutualFollowKey:  f.MutualFollowKey,
		Message:          "room follow completed",
	}
}

type DeleteFollow struct {
	Message  string `json:"message"`
}

func ToDeleteFollow() *DeleteFollow {
	return &DeleteFollow{
		Message: "delete follow completed",
	}
}
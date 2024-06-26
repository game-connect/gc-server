package model_test

import (
	"time"
	"testing"

	"github.com/game-connect/gc-server/domain/model"
)

func TestChatModel_EmptyRoom(t *testing.T) {
	tests := []struct {
		name     string
		chat     *model.Chat
		expected bool
	}{
		{
			name:     "Empty Chat",
			chat:     model.EmptyChat(),
			expected: true,
		},
		{
			name: "Not Empty Chat",
			chat: &model.Chat{
				ID:         1,
				ChatKey:    "test_key",
				ChannelKey: "test_key",
				UserKey:    "test_key",
				UserName:   "test_name",
				Content:    "content",
				ImagePath:  "/chat",
				PostedAt:    time.Now(), 
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.chat.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}

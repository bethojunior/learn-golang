package entities

import (
	"github.com/pborman/uuid"

	types "api/src/types/public"
)

func NewTweet(id string, description string) *types.Tweet {
	return &types.Tweet{
		ID: uuid.New(),
	}
}

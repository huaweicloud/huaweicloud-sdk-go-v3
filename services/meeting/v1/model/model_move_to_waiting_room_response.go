package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveToWaitingRoomResponse Response Object
type MoveToWaitingRoomResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MoveToWaitingRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveToWaitingRoomResponse struct{}"
	}

	return strings.Join([]string{"MoveToWaitingRoomResponse", string(data)}, " ")
}

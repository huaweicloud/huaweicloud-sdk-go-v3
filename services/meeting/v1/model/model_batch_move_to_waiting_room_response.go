package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMoveToWaitingRoomResponse Response Object
type BatchMoveToWaitingRoomResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchMoveToWaitingRoomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMoveToWaitingRoomResponse struct{}"
	}

	return strings.Join([]string{"BatchMoveToWaitingRoomResponse", string(data)}, " ")
}

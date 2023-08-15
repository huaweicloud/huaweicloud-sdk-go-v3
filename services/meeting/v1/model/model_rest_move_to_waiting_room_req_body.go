package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestMoveToWaitingRoomReqBody struct {

	// 与会者标识。
	ParticipantID string `json:"participantID"`
}

func (o RestMoveToWaitingRoomReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestMoveToWaitingRoomReqBody struct{}"
	}

	return strings.Join([]string{"RestMoveToWaitingRoomReqBody", string(data)}, " ")
}

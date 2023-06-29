package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BroadcastParticipantResponse Response Object
type BroadcastParticipantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BroadcastParticipantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BroadcastParticipantResponse struct{}"
	}

	return strings.Join([]string{"BroadcastParticipantResponse", string(data)}, " ")
}

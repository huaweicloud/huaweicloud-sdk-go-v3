package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowWaitingParticipantResponse Response Object
type AllowWaitingParticipantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AllowWaitingParticipantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowWaitingParticipantResponse struct{}"
	}

	return strings.Join([]string{"AllowWaitingParticipantResponse", string(data)}, " ")
}

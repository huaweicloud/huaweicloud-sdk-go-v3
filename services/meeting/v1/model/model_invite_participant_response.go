package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type InviteParticipantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InviteParticipantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteParticipantResponse struct{}"
	}

	return strings.Join([]string{"InviteParticipantResponse", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type MuteParticipantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MuteParticipantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MuteParticipantResponse struct{}"
	}

	return strings.Join([]string{"MuteParticipantResponse", string(data)}, " ")
}

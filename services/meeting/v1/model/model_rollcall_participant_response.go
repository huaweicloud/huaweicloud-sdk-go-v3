package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RollcallParticipantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RollcallParticipantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollcallParticipantResponse struct{}"
	}

	return strings.Join([]string{"RollcallParticipantResponse", string(data)}, " ")
}

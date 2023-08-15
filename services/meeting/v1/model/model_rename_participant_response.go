package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenameParticipantResponse Response Object
type RenameParticipantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RenameParticipantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameParticipantResponse struct{}"
	}

	return strings.Join([]string{"RenameParticipantResponse", string(data)}, " ")
}

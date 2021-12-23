package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchInviteMembersToChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchInviteMembersToChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInviteMembersToChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchInviteMembersToChannelResponse", string(data)}, " ")
}

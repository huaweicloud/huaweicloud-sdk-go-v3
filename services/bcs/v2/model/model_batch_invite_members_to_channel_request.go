package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInviteMembersToChannelRequest Request Object
type BatchInviteMembersToChannelRequest struct {
	Body *BatchInviteMembersToChannelRequestBody `json:"body,omitempty"`
}

func (o BatchInviteMembersToChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInviteMembersToChannelRequest struct{}"
	}

	return strings.Join([]string{"BatchInviteMembersToChannelRequest", string(data)}, " ")
}

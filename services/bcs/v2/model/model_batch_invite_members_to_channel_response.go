package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInviteMembersToChannelResponse Response Object
type BatchInviteMembersToChannelResponse struct {

	// 请求成功的结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchInviteMembersToChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInviteMembersToChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchInviteMembersToChannelResponse", string(data)}, " ")
}

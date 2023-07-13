package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMemberInviteRequest Request Object
type DeleteMemberInviteRequest struct {
	Body *DeleteMemberInviteRequestBody `json:"body,omitempty"`
}

func (o DeleteMemberInviteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberInviteRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberInviteRequest", string(data)}, " ")
}

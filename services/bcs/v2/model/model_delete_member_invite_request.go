package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteMemberInviteRequest struct {
	Body *DeleteMemberInviteRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DeleteMemberInviteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberInviteRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberInviteRequest", string(data)}, " ")
}

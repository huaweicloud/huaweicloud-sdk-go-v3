package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddServerGroupMemberRequest Request Object
type BatchAddServerGroupMemberRequest struct {
	ServerGroupId string `json:"server_group_id"`

	Body *BatchAddServerGroupMemberReq `json:"body,omitempty"`
}

func (o BatchAddServerGroupMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerGroupMemberRequest struct{}"
	}

	return strings.Join([]string{"BatchAddServerGroupMemberRequest", string(data)}, " ")
}

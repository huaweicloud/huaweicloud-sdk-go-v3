package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerGroupMemberRequest Request Object
type BatchDeleteServerGroupMemberRequest struct {
	ServerGroupId string `json:"server_group_id"`

	Body *BatchDeleteServerGroupMemberReq `json:"body,omitempty"`
}

func (o BatchDeleteServerGroupMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerGroupMemberRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerGroupMemberRequest", string(data)}, " ")
}

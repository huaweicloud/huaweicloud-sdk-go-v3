package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerGroupTagsRequest Request Object
type CreateServerGroupTagsRequest struct {

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`

	Body *CreateResourceTagReq `json:"body,omitempty"`
}

func (o CreateServerGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateServerGroupTagsRequest", string(data)}, " ")
}

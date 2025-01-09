package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupTagsRequest Request Object
type DeleteServerGroupTagsRequest struct {

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`

	Body *DeleteResourceTagReq `json:"body,omitempty"`
}

func (o DeleteServerGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupTagsRequest", string(data)}, " ")
}

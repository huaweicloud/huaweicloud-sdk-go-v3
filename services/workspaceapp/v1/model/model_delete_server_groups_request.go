package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupsRequest Request Object
type DeleteServerGroupsRequest struct {

	// 服务器组唯一标识
	ServerGroupId string `json:"server_group_id"`
}

func (o DeleteServerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupsRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupsRequest", string(data)}, " ")
}

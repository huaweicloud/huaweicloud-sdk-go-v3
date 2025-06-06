package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerGroupRequest Request Object
type ShowServerGroupRequest struct {

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`
}

func (o ShowServerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowServerGroupRequest", string(data)}, " ")
}

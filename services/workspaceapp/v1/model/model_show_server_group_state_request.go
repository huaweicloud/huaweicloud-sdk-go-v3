package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerGroupStateRequest Request Object
type ShowServerGroupStateRequest struct {

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`
}

func (o ShowServerGroupStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupStateRequest struct{}"
	}

	return strings.Join([]string{"ShowServerGroupStateRequest", string(data)}, " ")
}

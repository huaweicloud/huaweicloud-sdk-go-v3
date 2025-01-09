package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerGroupTagRequest Request Object
type ShowServerGroupTagRequest struct {

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`
}

func (o ShowServerGroupTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerGroupTagRequest struct{}"
	}

	return strings.Join([]string{"ShowServerGroupTagRequest", string(data)}, " ")
}

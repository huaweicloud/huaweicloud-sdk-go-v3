package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerGroupRequest Request Object
type UpdateServerGroupRequest struct {

	// 服务器组唯一标识
	ServerGroupId string `json:"server_group_id"`

	Body *UpdateServerGroupReq `json:"body,omitempty"`
}

func (o UpdateServerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerGroupRequest", string(data)}, " ")
}

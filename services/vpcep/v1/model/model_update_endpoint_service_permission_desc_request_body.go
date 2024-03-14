package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointServicePermissionDescRequestBody 更新终端节点服务白名单描述
type UpdateEndpointServicePermissionDescRequestBody struct {
	Permission *EpsUpdatePermissionDesc `json:"permission"`
}

func (o UpdateEndpointServicePermissionDescRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServicePermissionDescRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServicePermissionDescRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePermissionDescRequest 更新终端节点服务白名单描述
type UpdatePermissionDescRequest struct {
	Permission *EpsUpdatePermissionDesc `json:"permission"`
}

func (o UpdatePermissionDescRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionDescRequest struct{}"
	}

	return strings.Join([]string{"UpdatePermissionDescRequest", string(data)}, " ")
}

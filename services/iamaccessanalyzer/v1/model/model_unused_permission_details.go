package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnusedPermissionDetails struct {

	// 权限对应的云服务名称。
	Service string `json:"service"`

	// 未使用的操作列表。
	Actions []UnusedAction `json:"actions"`
}

func (o UnusedPermissionDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedPermissionDetails struct{}"
	}

	return strings.Join([]string{"UnusedPermissionDetails", string(data)}, " ")
}

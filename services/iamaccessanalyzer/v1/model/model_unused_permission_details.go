package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnusedPermissionDetails 未使用权限分析详细结果。
type UnusedPermissionDetails struct {

	// 权限对应的云服务名称。
	Service string `json:"service"`

	// 用户使用云服务的最后访问时间。
	LastAccessed *sdktime.SdkTime `json:"last_accessed,omitempty"`

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

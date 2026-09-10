package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcePackageInfo struct {

	// 资源包ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 引擎名称。
	EngineName *string `json:"engine_name,omitempty"`

	// 已使用配额。
	UsedQuota *int32 `json:"used_quota,omitempty"`

	// 总配额。
	TotalQuota *int32 `json:"total_quota,omitempty"`

	// 资源包状态。
	Status *string `json:"status,omitempty"`
}

func (o ResourcePackageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePackageInfo struct{}"
	}

	return strings.Join([]string{"ResourcePackageInfo", string(data)}, " ")
}

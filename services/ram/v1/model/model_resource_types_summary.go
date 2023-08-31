package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTypesSummary 对接RAM云服务的资源类型和区域等信息。
type ResourceTypesSummary struct {

	// 资源所在区域名称。
	RegionId string `json:"region_id"`

	// 资源类型名称。
	ResourceType string `json:"resource_type"`
}

func (o ResourceTypesSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTypesSummary struct{}"
	}

	return strings.Join([]string{"ResourceTypesSummary", string(data)}, " ")
}

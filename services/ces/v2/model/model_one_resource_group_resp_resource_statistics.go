package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneResourceGroupRespResourceStatistics 资源数量统计信息
type OneResourceGroupRespResourceStatistics struct {

	// 告警中的资源数
	Unhealthy *int32 `json:"unhealthy,omitempty"`

	// 资源总数
	Total *int32 `json:"total,omitempty"`

	// 已触发的资源数
	EventUnhealthy *int32 `json:"event_unhealthy,omitempty"`

	// 资源类型数
	Namespaces *int32 `json:"namespaces,omitempty"`
}

func (o OneResourceGroupRespResourceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneResourceGroupRespResourceStatistics struct{}"
	}

	return strings.Join([]string{"OneResourceGroupRespResourceStatistics", string(data)}, " ")
}

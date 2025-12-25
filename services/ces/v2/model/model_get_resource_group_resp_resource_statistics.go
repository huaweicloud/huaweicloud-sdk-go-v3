package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceGroupRespResourceStatistics 资源数量统计信息
type GetResourceGroupRespResourceStatistics struct {

	// **参数解释** 告警中的资源数。 **取值范围** 0-9999999
	Unhealthy *int32 `json:"unhealthy,omitempty"`

	// **参数解释** 资源总数。 **取值范围** 0-9999999
	Total *int32 `json:"total,omitempty"`

	// **参数解释** 已触发的资源数。 **取值范围** 0-9999999
	EventUnhealthy *int32 `json:"event_unhealthy,omitempty"`

	// **参数解释** 资源类型数。 **取值范围** 0-9999999
	Namespaces *int32 `json:"namespaces,omitempty"`
}

func (o GetResourceGroupRespResourceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceGroupRespResourceStatistics struct{}"
	}

	return strings.Join([]string{"GetResourceGroupRespResourceStatistics", string(data)}, " ")
}

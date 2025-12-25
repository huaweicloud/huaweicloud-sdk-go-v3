package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneResourceGroupRespResourceStatistics **参数解释** 资源分组中的资源信息统计。
type OneResourceGroupRespResourceStatistics struct {

	// **参数解释** 该资源分组中当前处在告警状态的资源个数。  **取值范围** 在[0,9999999]区间内。
	Unhealthy *int32 `json:"unhealthy,omitempty"`

	// **参数解释** 该资源分组中资源的总个数。  **取值范围** 在[0,9999999]区间内。
	Total *int32 `json:"total,omitempty"`

	// **参数解释** 该资源分组中已触发的资源个数。  **取值范围** 在[0,9999999]区间内。
	EventUnhealthy *int32 `json:"event_unhealthy,omitempty"`

	// **参数解释** 该资源分组中选择的资源类型个数，如资源分组添加了弹性云服务、弹性公网IP和带宽则值为2。 **取值范围** 在[0,9999999]区间内。
	Namespaces *int32 `json:"namespaces,omitempty"`
}

func (o OneResourceGroupRespResourceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneResourceGroupRespResourceStatistics struct{}"
	}

	return strings.Join([]string{"OneResourceGroupRespResourceStatistics", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源分组中的资源信息统计。
type InstanceStatistics struct {
	// 该资源分组中当前处在告警状态的资源个数。

	Unhealth *int32 `json:"unhealth,omitempty"`
	// 该资源分组中资源的总个数。

	Total *int32 `json:"total,omitempty"`
	// 该资源分组中选择的资源类型个数，如资源分组添加了弹性云服务、弹性公网IP和带宽则值为2。

	TypeStatistics *int32 `json:"type_statistics,omitempty"`
}

func (o InstanceStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatistics struct{}"
	}

	return strings.Join([]string{"InstanceStatistics", string(data)}, " ")
}

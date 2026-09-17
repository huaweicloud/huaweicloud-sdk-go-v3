package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpaceTrend SpaceTrend对象
type SpaceTrend struct {

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 指标名
	Name *string `json:"name,omitempty"`

	// 指标值列表
	Series *[]float64 `json:"series,omitempty"`

	// 时间戳列表
	Timestamps *[]int64 `json:"timestamps,omitempty"`
}

func (o SpaceTrend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpaceTrend struct{}"
	}

	return strings.Join([]string{"SpaceTrend", string(data)}, " ")
}

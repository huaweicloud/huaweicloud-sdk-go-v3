package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogPoint 节点趋势点
type SlowLogPoint struct {

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// 趋势数量列表
	TrendData *[]SlowLogTrendPoint `json:"trend_data,omitempty"`
}

func (o SlowLogPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogPoint struct{}"
	}

	return strings.Join([]string{"SlowLogPoint", string(data)}, " ")
}

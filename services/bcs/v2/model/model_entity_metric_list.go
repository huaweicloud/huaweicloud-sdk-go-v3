package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BCS组织维度监控数据列表结构
type EntityMetricList struct {
	// 指标对象列表。

	Dimensions *[]Dimension `json:"dimensions,omitempty"`
	// 监控数据列表项目。

	Values *[]EntityMetricListItem `json:"values,omitempty"`
}

func (o EntityMetricList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityMetricList struct{}"
	}

	return strings.Join([]string{"EntityMetricList", string(data)}, " ")
}

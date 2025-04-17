package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HierarchicalValue 多层级告警阈值。当hierarchical_value和value同时使用时以hierarchical_value为准。 创建或修改告警规则以下2种场景只支持设置一个阈值：   1.告警类型为`指标告警`且告警策略为`所有策略都满足才告警`的场景。   2.告警类型为`事件告警`的场景。
type HierarchicalValue struct {

	// 紧急级别的阈值
	Critical *float64 `json:"critical,omitempty"`

	// 重要级别的阈值
	Major *float64 `json:"major,omitempty"`

	// 次要级别的阈值
	Minor *float64 `json:"minor,omitempty"`

	// 提示级别的阈值
	Info *float64 `json:"info,omitempty"`
}

func (o HierarchicalValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HierarchicalValue struct{}"
	}

	return strings.Join([]string{"HierarchicalValue", string(data)}, " ")
}

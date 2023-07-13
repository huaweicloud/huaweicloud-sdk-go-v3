package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopoLine 组件之间调用指向线。
type TopoLine struct {

	// 开始节点。
	FromNode *string `json:"from_node,omitempty"`

	// 结束节点。
	ToNode *string `json:"to_node,omitempty"`

	// 指向。
	Direction *string `json:"direction,omitempty"`

	// 采集器名称。
	Collector *string `json:"collector,omitempty"`

	// 环境id。
	TargetEnvId *int64 `json:"target_env_id,omitempty"`

	// 线条上的提示信息。
	Hints map[string]string `json:"hints,omitempty"`

	// 过滤值。
	FilterValue *string `json:"filter_value,omitempty"`
}

func (o TopoLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopoLine struct{}"
	}

	return strings.Join([]string{"TopoLine", string(data)}, " ")
}

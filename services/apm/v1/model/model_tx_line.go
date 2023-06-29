package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TxLine 组件之间调用指向线。
type TxLine struct {

	// 开始节点。
	TxFromNode *string `json:"tx_from_node,omitempty"`

	// 结束节点。
	TxToNode *string `json:"tx_to_node,omitempty"`

	// 调用次数。
	InvokeCount *int64 `json:"invoke_count,omitempty"`

	// 平均响应时间。
	Rt *float64 `json:"rt,omitempty"`

	// 错误数。
	ErrorCount *int64 `json:"error_count,omitempty"`

	// 类型。
	Type *string `json:"type,omitempty"`

	// 指向。
	Direction *string `json:"direction,omitempty"`
}

func (o TxLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TxLine struct{}"
	}

	return strings.Join([]string{"TxLine", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertExpression struct {

	// 表达式操作符
	ExpressionOperator *string `json:"expression_operator,omitempty"`

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 指标操作符
	MetricOperator *string `json:"metric_operator,omitempty"`

	// 指标阈值
	MetricThreshold *int32 `json:"metric_threshold,omitempty"`
}

func (o AlertExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertExpression struct{}"
	}

	return strings.Join([]string{"AlertExpression", string(data)}, " ")
}

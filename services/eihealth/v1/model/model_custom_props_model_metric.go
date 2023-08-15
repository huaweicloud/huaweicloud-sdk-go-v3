package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomPropsModelMetric 自定义属性建模的评估指标
type CustomPropsModelMetric struct {

	// 评估指标的名称
	Name *string `json:"name,omitempty"`

	// 评估指标的评估结果
	Value *float32 `json:"value,omitempty"`
}

func (o CustomPropsModelMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPropsModelMetric struct{}"
	}

	return strings.Join([]string{"CustomPropsModelMetric", string(data)}, " ")
}

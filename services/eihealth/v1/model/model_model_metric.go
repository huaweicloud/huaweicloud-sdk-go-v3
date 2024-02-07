package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelMetric 模型评估指标
type ModelMetric struct {

	// 评估指标的名称
	Name *string `json:"name,omitempty"`

	// 评估指标的评估结果
	Value *float32 `json:"value,omitempty"`
}

func (o ModelMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelMetric struct{}"
	}

	return strings.Join([]string{"ModelMetric", string(data)}, " ")
}

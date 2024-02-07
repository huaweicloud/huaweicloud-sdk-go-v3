package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelDtoModelMetric 模型评估指标
type ModelDtoModelMetric struct {

	// 评估指标的名称
	Name *string `json:"name,omitempty"`

	// 评估指标的评估结果
	Value *float32 `json:"value,omitempty"`
}

func (o ModelDtoModelMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelDtoModelMetric struct{}"
	}

	return strings.Join([]string{"ModelDtoModelMetric", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 点位处理配置结构体
type ProcessingConfigDto struct {
	Validity *PointValidityingDto `json:"validity,omitempty"`

	// 点位流公式配置字段
	StreamFormula *string `json:"stream_formula,omitempty"`

	Scaling *PointScalingDto `json:"scaling,omitempty"`

	Clean *PointCleanDto `json:"clean,omitempty"`
}

func (o ProcessingConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessingConfigDto struct{}"
	}

	return strings.Join([]string{"ProcessingConfigDto", string(data)}, " ")
}

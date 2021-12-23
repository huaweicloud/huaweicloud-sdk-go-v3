package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 维度。
type DimensionSeries struct {
	// 名称。

	Name *string `json:"name,omitempty"`
	// 具体数值。

	Value *string `json:"value,omitempty"`
}

func (o DimensionSeries) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionSeries struct{}"
	}

	return strings.Join([]string{"DimensionSeries", string(data)}, " ")
}

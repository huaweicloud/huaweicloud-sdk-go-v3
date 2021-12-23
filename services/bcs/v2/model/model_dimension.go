package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指标维度。
type Dimension struct {
	// 维度名称。

	Name *string `json:"name,omitempty"`
	// 维度取值。

	Value *string `json:"value,omitempty"`
}

func (o Dimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dimension struct{}"
	}

	return strings.Join([]string{"Dimension", string(data)}, " ")
}

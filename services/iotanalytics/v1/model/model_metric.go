package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Metric struct {

	// 指标名称
	Name *string `json:"name,omitempty"`

	// 指标值类型
	Type *string `json:"type,omitempty"`

	// 指标描述
	Description *string `json:"description,omitempty"`

	// 指标单位
	Unit *string `json:"unit,omitempty"`
}

func (o Metric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metric struct{}"
	}

	return strings.Join([]string{"Metric", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceUsage struct {

	// 当前使用量。
	Value *float64 `json:"value,omitempty"`

	// 总量。
	Total *float64 `json:"total,omitempty"`

	// 对比值。
	Contrast *float64 `json:"contrast,omitempty"`

	// 单位。
	Unit *string `json:"unit,omitempty"`
}

func (o ResourceUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceUsage struct{}"
	}

	return strings.Join([]string{"ResourceUsage", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomBaselineRulePatchItem 自定义基线补丁
type CustomBaselineRulePatchItem struct {

	// 补丁名称
	Name *string `json:"name,omitempty"`

	// 补丁版本
	Version *string `json:"version,omitempty"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`
}

func (o CustomBaselineRulePatchItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomBaselineRulePatchItem struct{}"
	}

	return strings.Join([]string{"CustomBaselineRulePatchItem", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义字段
type IssueCustomField struct {

	// 自定义字段
	Name *string `json:"name,omitempty"`

	// 自定义字段
	CustomField *string `json:"custom_field,omitempty"`

	// 自定义字段的可选值，多个值以英文逗号区分
	Options *string `json:"options,omitempty"`

	// 自定义字段类型
	Type *string `json:"type,omitempty"`

	// 自定义字段支持的工作项类型 2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
	TrackerIds *[]int32 `json:"tracker_ids,omitempty"`
}

func (o IssueCustomField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCustomField struct{}"
	}

	return strings.Join([]string{"IssueCustomField", string(data)}, " ")
}

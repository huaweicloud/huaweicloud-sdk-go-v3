package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSpecsRequest Request Object
type ListEventSpecsRequest struct {

	// 事件配置名称
	SpecName *string `json:"spec_name,omitempty"`

	// 事件类别
	Category *string `json:"category,omitempty"`

	// 事件级别
	Severity *string `json:"severity,omitempty"`

	// 事件源类别
	SourceType *string `json:"source_type,omitempty"`

	// 事件标签
	Tag *string `json:"tag,omitempty"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 限制条目数
	Limit *string `json:"limit,omitempty"`
}

func (o ListEventSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSpecsRequest struct{}"
	}

	return strings.Join([]string{"ListEventSpecsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowListOfEventSampleRequest Request Object
type ShowListOfEventSampleRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询的事件示例名称，模糊匹配
	Name *string `json:"name,omitempty"`

	// 指定查询事件示例的事件类型名称
	EventTypeName *string `json:"event_type_name,omitempty"`

	// 指定查询事件示例的事件源
	EventSourceId *string `json:"event_source_id,omitempty"`
}

func (o ShowListOfEventSampleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListOfEventSampleRequest struct{}"
	}

	return strings.Join([]string{"ShowListOfEventSampleRequest", string(data)}, " ")
}

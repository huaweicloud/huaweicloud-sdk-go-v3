package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessControlTaskRequest Request Object
type ListAccessControlTaskRequest struct {

	// 查询起始时间戳（毫秒），不传默认为当前时间，需与结束时间戳同时指定，时间跨度不能超过7天。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询结束时间戳（毫秒），不传默认3天前，需与开始时间戳同时指定，时间跨度不能超过7天。
	EndTime *int64 `json:"end_time,omitempty"`

	// 偏移量：特定数据字段与起始数据字段位置的距离，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询数据条数，上限为10000，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 任务状态，状态类型：processing：处理中；succeed：完成；failed：失败。
	Status *string `json:"status,omitempty"`
}

func (o ListAccessControlTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessControlTaskRequest struct{}"
	}

	return strings.Join([]string{"ListAccessControlTaskRequest", string(data)}, " ")
}

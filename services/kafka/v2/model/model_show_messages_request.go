package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMessagesRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// Topic名称。

	Topic string `json:"topic"`
	// 查询起始时间，为unix时间戳格式，默认值为0。

	StartTime *string `json:"start_time,omitempty"`
	// 查询结束时间，为unix时间戳格式，默认值为系统当前时间。

	EndTime *string `json:"end_time,omitempty"`
	// 单页返回消息数，默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0。

	Offset *int32 `json:"offset,omitempty"`
	// 分区编号，默认值为-1，若传入值为-1，则查询所有分区。

	Partition *string `json:"partition,omitempty"`
}

func (o ShowMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessagesRequest struct{}"
	}

	return strings.Join([]string{"ShowMessagesRequest", string(data)}, " ")
}

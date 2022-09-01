package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowGroupsRespGroupGroupMessageOffsets struct {

	// 分区编号。
	Partition *int32 `json:"partition,omitempty" xml:"partition"`

	// 剩余可消费消息数，即消息堆积数。
	Lag *int32 `json:"lag,omitempty" xml:"lag"`

	// topic名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 当前消费进度。
	MessageCurrentOffset *int32 `json:"message_current_offset,omitempty" xml:"message_current_offset"`

	// 最大消息位置（LEO）。
	MessageLogEndOffset *int32 `json:"message_log_end_offset,omitempty" xml:"message_log_end_offset"`
}

func (o ShowGroupsRespGroupGroupMessageOffsets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupGroupMessageOffsets struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupGroupMessageOffsets", string(data)}, " ")
}

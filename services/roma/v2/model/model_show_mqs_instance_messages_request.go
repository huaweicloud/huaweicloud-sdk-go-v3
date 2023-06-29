package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMqsInstanceMessagesRequest Request Object
type ShowMqsInstanceMessagesRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// topic名称。
	Topic string `json:"topic"`

	// 是否按照时间排序。
	Asc *bool `json:"asc,omitempty"`

	// 开始时间。Unix毫秒时间戳。 查询消息偏移量时，为必选参数。
	StartTime string `json:"start_time"`

	// 结束时间。Unix毫秒时间戳。 查询消息偏移量时，为必选参数。
	EndTime string `json:"end_time"`

	// 查询消息的数量。
	Limit *string `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *string `json:"offset,omitempty"`

	// 是否下载。
	Download *bool `json:"download,omitempty"`

	// 消息偏移量。 查询消息内容时，为必选参数。 若start_time、end_time参数不为空，该参数无效。
	MessageOffset *string `json:"message_offset,omitempty"`

	// 分区。 查询消息内容时，为必选参数。 若start_time、end_time参数不为空，该参数无效。
	Partition *string `json:"partition,omitempty"`

	// 消息key。
	Key *string `json:"key,omitempty"`

	// 消息ID。
	MessageId *string `json:"message_id,omitempty"`

	// 消息标签。
	Tag *string `json:"tag,omitempty"`
}

func (o ShowMqsInstanceMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceMessagesRequest struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceMessagesRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPartitionBeginningMessageResponse struct {

	// Topic名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 分区编号。
	Partition *int32 `json:"partition,omitempty" xml:"partition"`

	// 最新消息位置。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 生产消息的时间。 格式为Unix时间戳。单位为毫秒。
	Timestamp      *int64 `json:"timestamp,omitempty" xml:"timestamp"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPartitionBeginningMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionBeginningMessageResponse struct{}"
	}

	return strings.Join([]string{"ShowPartitionBeginningMessageResponse", string(data)}, " ")
}

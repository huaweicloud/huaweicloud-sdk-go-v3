package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerGroupInfo struct {

	// 创建的消费组名
	ConsumerGroupName string `json:"consumer_group_name"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 日志组ID
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 超时时间，单位秒
	Timeout *int32 `json:"timeout,omitempty"`
}

func (o ConsumerGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroupInfo struct{}"
	}

	return strings.Join([]string{"ConsumerGroupInfo", string(data)}, " ")
}

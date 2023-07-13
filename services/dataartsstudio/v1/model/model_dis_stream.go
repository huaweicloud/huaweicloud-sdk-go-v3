package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisStream struct {

	// 通道名称
	StreamName *string `json:"stream_name,omitempty"`

	// 通道GUID
	StreamGuid *string `json:"stream_guid,omitempty"`

	// 通道的唯一标识名称
	StreamQualifiedName *string `json:"stream_qualified_name,omitempty"`

	// 分区数
	PartitionCount *int32 `json:"partition_count,omitempty"`

	// dis的app数目
	AppCount *int32 `json:"app_count,omitempty"`

	// 转储任务数
	TaskCount *int32 `json:"task_count,omitempty"`
}

func (o DisStream) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisStream struct{}"
	}

	return strings.Join([]string{"DisStream", string(data)}, " ")
}

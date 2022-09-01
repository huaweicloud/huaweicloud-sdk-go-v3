package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowInstanceTopicDetailRespPartitions struct {

	// 分区ID。
	Partition *int32 `json:"partition,omitempty" xml:"partition"`

	// leader副本所在节点的id。
	Leader *int32 `json:"leader,omitempty" xml:"leader"`

	// 分区leader副本的LEO（Log End Offset）。
	Leo *int32 `json:"leo,omitempty" xml:"leo"`

	// 分区高水位（HW，High Watermark）。
	Hw *int32 `json:"hw,omitempty" xml:"hw"`

	// 分区leader副本的LSO（Log Start Offset）。
	Lso *int32 `json:"lso,omitempty" xml:"lso"`

	// 分区上次写入消息的时间。  格式为Unix时间戳。  单位：毫秒。
	LastUpdateTimestamp *int64 `json:"last_update_timestamp,omitempty" xml:"last_update_timestamp"`

	// 副本列表。
	Replicas *[]ShowInstanceTopicDetailRespReplicas `json:"replicas,omitempty" xml:"replicas"`
}

func (o ShowInstanceTopicDetailRespPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailRespPartitions struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailRespPartitions", string(data)}, " ")
}

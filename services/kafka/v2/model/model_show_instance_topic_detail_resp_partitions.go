package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowInstanceTopicDetailRespPartitions struct {

	// **参数解释**： 分区ID。 **取值范围**： 不涉及
	Partition *int32 `json:"partition,omitempty"`

	// **参数解释**： leader副本所在节点的id。 **取值范围**： 不涉及
	Leader *int32 `json:"leader,omitempty"`

	// **参数解释**： 分区leader副本的LEO（Log End Offset）。 **取值范围**： 不涉及
	Leo *int32 `json:"leo,omitempty"`

	// **参数解释**： 分区高水位（HW，High Watermark）。 **取值范围**： 不涉及
	Hw *int32 `json:"hw,omitempty"`

	// **参数解释**： 分区leader副本的LSO（Log Start Offset）。 **取值范围**： 不涉及
	Lso *int32 `json:"lso,omitempty"`

	// **参数解释**： 分区上次写入消息的时间。  格式为Unix时间戳。  单位：毫秒。 **取值范围**： 不涉及
	LastUpdateTimestamp *int64 `json:"last_update_timestamp,omitempty"`

	// **参数解释**： 副本列表。
	Replicas *[]ShowInstanceTopicDetailRespReplicas `json:"replicas,omitempty"`
}

func (o ShowInstanceTopicDetailRespPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailRespPartitions struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailRespPartitions", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceTopicsResponse Response Object
type ListInstanceTopicsResponse struct {

	// **参数解释**： Topic总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 分页查询的大小。 **取值范围**： 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**： 剩余分区数。 **取值范围**： 不涉及。
	RemainPartitions *int32 `json:"remain_partitions,omitempty"`

	// **参数解释**： 分区总数。 **取值范围**： 不涉及。
	MaxPartitions *int32 `json:"max_partitions,omitempty"`

	// **参数解释**： 单个Topic最大占用分区数。 **取值范围**： 不涉及。
	TopicMaxPartitions *int32 `json:"topic_max_partitions,omitempty"`

	// **参数解释**： Topic列表。
	Topics         *[]TopicEntity `json:"topics,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListInstanceTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowInstanceTopicDetailRespReplicas struct {

	// **参数解释**： 副本所在的节点ID。 **取值范围**： 不涉及
	Broker *int32 `json:"broker,omitempty"`

	// **参数解释**： 该副本是否为leader。 **取值范围**： - true：该副本是leader。 - false：该副本不是leader。
	Leader *bool `json:"leader,omitempty"`

	// **参数解释**： 该副本是否在ISR副本中。 **取值范围**： - true：在ISR副本中。 - false：不在ISR副本中。
	InSync *bool `json:"in_sync,omitempty"`

	// **参数解释**： 该副本当前日志大小。单位：Byte。 **取值范围**： 不涉及
	Size *int32 `json:"size,omitempty"`

	// **参数解释**： 该副本当前落后hw的消息数。 **取值范围**： 不涉及
	Lag *int64 `json:"lag,omitempty"`
}

func (o ShowInstanceTopicDetailRespReplicas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailRespReplicas struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailRespReplicas", string(data)}, " ")
}

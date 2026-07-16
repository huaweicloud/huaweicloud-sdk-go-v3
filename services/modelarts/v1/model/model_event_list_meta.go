package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventListMeta 事件列表元信息。
type EventListMeta struct {

	// **参数描述**：分页标记。 **取值范围**：不涉及。
	Continue *string `json:"continue,omitempty"`

	// **参数描述**：分页剩余数量。 **取值范围**：不涉及。
	RemainingItemCount *int32 `json:"remainingItemCount,omitempty"`
}

func (o EventListMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventListMeta struct{}"
	}

	return strings.Join([]string{"EventListMeta", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionMemoryContextInfoResult struct {

	// **参数解释**: 内存上下文名称。 **取值范围**: 不涉及。
	ContextName *string `json:"context_name,omitempty"`

	// **参数解释**: 会话上下文数量。 **取值范围**: 大于等于0。
	Amount *int32 `json:"amount,omitempty"`

	// **参数解释**: 会话上下文总大小。 **取值范围**: 大于等于0。
	Size *float64 `json:"size,omitempty"`
}

func (o SessionMemoryContextInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionMemoryContextInfoResult struct{}"
	}

	return strings.Join([]string{"SessionMemoryContextInfoResult", string(data)}, " ")
}

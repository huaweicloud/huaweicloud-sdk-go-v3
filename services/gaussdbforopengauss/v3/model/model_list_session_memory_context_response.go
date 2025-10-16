package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionMemoryContextResponse Response Object
type ListSessionMemoryContextResponse struct {

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 会话内存上下文列表。
	MemoryContextInfo *[]SessionMemoryContextInfoResult `json:"memory_context_info,omitempty"`
	HttpStatusCode    int                               `json:"-"`
}

func (o ListSessionMemoryContextResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionMemoryContextResponse struct{}"
	}

	return strings.Join([]string{"ListSessionMemoryContextResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWaitEventResponse Response Object
type ListWaitEventResponse struct {

	// **参数解释**: 等待事件的总数量。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	Rows           *WaitEventResult `json:"rows,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListWaitEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWaitEventResponse struct{}"
	}

	return strings.Join([]string{"ListWaitEventResponse", string(data)}, " ")
}

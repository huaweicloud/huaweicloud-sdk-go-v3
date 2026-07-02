package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectRealtimeSessionResponse Response Object
type CollectRealtimeSessionResponse struct {

	// **参数解释**：  收集全部实时会话信息的任务ID。  **取值范围**：  不涉及。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectRealtimeSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectRealtimeSessionResponse struct{}"
	}

	return strings.Join([]string{"CollectRealtimeSessionResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncResourceAgentResponse Response Object
type SyncResourceAgentResponse struct {

	// **参数解释：** 任务id。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncResourceAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourceAgentResponse struct{}"
	}

	return strings.Join([]string{"SyncResourceAgentResponse", string(data)}, " ")
}

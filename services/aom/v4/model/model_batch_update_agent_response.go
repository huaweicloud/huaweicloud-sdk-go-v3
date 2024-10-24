package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAgentResponse Response Object
type BatchUpdateAgentResponse struct {

	// 执行状态： - true：成功下发任务。 - false：失败下发任务。
	State          *bool `json:"state,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o BatchUpdateAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAgentResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateAgentResponse", string(data)}, " ")
}

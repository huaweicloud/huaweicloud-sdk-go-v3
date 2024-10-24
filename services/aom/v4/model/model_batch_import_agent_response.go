package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportAgentResponse Response Object
type BatchImportAgentResponse struct {

	// 执行状态： - true：成功下发任务。 - false：失败下发任务。
	State          *bool `json:"state,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o BatchImportAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportAgentResponse struct{}"
	}

	return strings.Join([]string{"BatchImportAgentResponse", string(data)}, " ")
}

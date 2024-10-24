package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportAgentRequest Request Object
type BatchImportAgentRequest struct {

	// region id，例如：cn-north-7
	Region string `json:"region"`

	Body *AgentBatchImportParamNew `json:"body,omitempty"`
}

func (o BatchImportAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportAgentRequest struct{}"
	}

	return strings.Join([]string{"BatchImportAgentRequest", string(data)}, " ")
}

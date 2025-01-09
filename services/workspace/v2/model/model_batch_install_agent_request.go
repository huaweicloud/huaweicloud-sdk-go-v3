package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInstallAgentRequest Request Object
type BatchInstallAgentRequest struct {
	Body *BatchInstallAgentReq `json:"body,omitempty"`
}

func (o BatchInstallAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInstallAgentRequest struct{}"
	}

	return strings.Join([]string{"BatchInstallAgentRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentConfigResponse Response Object
type ShowAgentConfigResponse struct {
	Code *string `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`

	Extend *interface{} `json:"extend,omitempty"`

	Result         *AgentConfig `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAgentConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentConfigResponse", string(data)}, " ")
}

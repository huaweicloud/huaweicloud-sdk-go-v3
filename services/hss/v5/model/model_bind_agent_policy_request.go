package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindAgentPolicyRequest Request Object
type BindAgentPolicyRequest struct {
	Body *BindAgentPolicyRequestInfo `json:"body,omitempty"`
}

func (o BindAgentPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindAgentPolicyRequest struct{}"
	}

	return strings.Join([]string{"BindAgentPolicyRequest", string(data)}, " ")
}

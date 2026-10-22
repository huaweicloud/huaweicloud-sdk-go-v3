package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiGuardPolicyRequest Request Object
type CreateAiGuardPolicyRequest struct {
	Body *CreateAiGuardPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateAiGuardPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiGuardPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAiGuardPolicyRequest", string(data)}, " ")
}

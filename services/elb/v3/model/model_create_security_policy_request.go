package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityPolicyRequest Request Object
type CreateSecurityPolicyRequest struct {
	Body *CreateSecurityPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyRequest", string(data)}, " ")
}

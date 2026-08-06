package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKeyPolicyRequest Request Object
type CreateKeyPolicyRequest struct {
	Body *CreateKeyPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateKeyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateKeyPolicyRequest", string(data)}, " ")
}

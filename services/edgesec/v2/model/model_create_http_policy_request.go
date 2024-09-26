package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpPolicyRequest Request Object
type CreateHttpPolicyRequest struct {
	Body *CreateHttpPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateHttpPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpPolicyRequest", string(data)}, " ")
}

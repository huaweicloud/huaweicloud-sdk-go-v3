package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateL7policyRequest struct {
	Body *CreateL7policyRequestBody `json:"body,omitempty"`
}

func (o CreateL7policyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7policyRequest struct{}"
	}

	return strings.Join([]string{"CreateL7policyRequest", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingPolicyRequest Request Object
type CreateScalingPolicyRequest struct {
	Body *CreateScalingPolicyOption `json:"body,omitempty"`
}

func (o CreateScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyRequest", string(data)}, " ")
}

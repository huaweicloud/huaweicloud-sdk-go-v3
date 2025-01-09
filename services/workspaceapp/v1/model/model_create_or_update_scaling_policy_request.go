package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateScalingPolicyRequest Request Object
type CreateOrUpdateScalingPolicyRequest struct {
	Body *CreateOrUpdateScalingPolicyReq `json:"body,omitempty"`
}

func (o CreateOrUpdateScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateScalingPolicyRequest", string(data)}, " ")
}

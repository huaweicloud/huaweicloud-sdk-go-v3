package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateScaleOutPolicyRequest struct {
	Body *CreateScaleOutPolicyReq `json:"body,omitempty"`
}

func (o CreateScaleOutPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScaleOutPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateScaleOutPolicyRequest", string(data)}, " ")
}

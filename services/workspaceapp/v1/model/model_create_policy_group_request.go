package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyGroupRequest Request Object
type CreatePolicyGroupRequest struct {
	Body *CreatePolicyGroupReq `json:"body,omitempty"`
}

func (o CreatePolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyGroupRequest", string(data)}, " ")
}

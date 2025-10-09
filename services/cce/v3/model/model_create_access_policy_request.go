package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessPolicyRequest Request Object
type CreateAccessPolicyRequest struct {
	Body *AccessPolicy `json:"body,omitempty"`
}

func (o CreateAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAccessPolicyRequest", string(data)}, " ")
}

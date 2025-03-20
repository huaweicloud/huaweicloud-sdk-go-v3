package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyVersionV5Response Response Object
type CreatePolicyVersionV5Response struct {
	PolicyVersion  *PolicyVersion `json:"policy_version,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreatePolicyVersionV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyVersionV5Response struct{}"
	}

	return strings.Join([]string{"CreatePolicyVersionV5Response", string(data)}, " ")
}

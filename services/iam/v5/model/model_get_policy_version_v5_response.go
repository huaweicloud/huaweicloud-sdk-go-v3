package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPolicyVersionV5Response Response Object
type GetPolicyVersionV5Response struct {
	PolicyVersion  *PolicyVersion `json:"policy_version,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o GetPolicyVersionV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyVersionV5Response struct{}"
	}

	return strings.Join([]string{"GetPolicyVersionV5Response", string(data)}, " ")
}

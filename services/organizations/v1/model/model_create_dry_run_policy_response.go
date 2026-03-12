package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDryRunPolicyResponse Response Object
type CreateDryRunPolicyResponse struct {
	Policy         *PolicyDto `json:"policy,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CreateDryRunPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDryRunPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateDryRunPolicyResponse", string(data)}, " ")
}

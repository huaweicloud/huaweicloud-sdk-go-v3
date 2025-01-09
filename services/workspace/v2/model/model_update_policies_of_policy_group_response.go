package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePoliciesOfPolicyGroupResponse Response Object
type UpdatePoliciesOfPolicyGroupResponse struct {
	Policies       *Policies `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePoliciesOfPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoliciesOfPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdatePoliciesOfPolicyGroupResponse", string(data)}, " ")
}

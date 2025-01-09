package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoliciesOfPolicyGroupResponse Response Object
type ListPoliciesOfPolicyGroupResponse struct {
	Policies       *Policies `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPoliciesOfPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesOfPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListPoliciesOfPolicyGroupResponse", string(data)}, " ")
}

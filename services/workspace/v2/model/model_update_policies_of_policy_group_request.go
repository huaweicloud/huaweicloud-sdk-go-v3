package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePoliciesOfPolicyGroupRequest Request Object
type UpdatePoliciesOfPolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`

	Body *ModifyPolicyRequest `json:"body,omitempty"`
}

func (o UpdatePoliciesOfPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoliciesOfPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdatePoliciesOfPolicyGroupRequest", string(data)}, " ")
}

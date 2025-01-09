package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyPolicyGroupRequest struct {
	PolicyGroup *PolicyGroupForUpdate `json:"policy_group,omitempty"`
}

func (o ModifyPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"ModifyPolicyGroupRequest", string(data)}, " ")
}

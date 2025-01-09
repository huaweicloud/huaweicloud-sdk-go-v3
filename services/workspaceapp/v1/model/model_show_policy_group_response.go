package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyGroupResponse Response Object
type ShowPolicyGroupResponse struct {
	PolicyGroup    *PolicyGroup `json:"policy_group,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyGroupResponse", string(data)}, " ")
}

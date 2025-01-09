package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyOfPolicyGroupResponse Response Object
type ListPolicyOfPolicyGroupResponse struct {
	Policies       *Policies `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPolicyOfPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyOfPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyOfPolicyGroupResponse", string(data)}, " ")
}

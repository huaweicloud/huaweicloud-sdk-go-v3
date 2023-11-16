package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGovernancePolicyByPolicyIdResponse Response Object
type ListGovernancePolicyByPolicyIdResponse struct {

	// 治理策略名称
	Name *string `json:"name,omitempty"`

	Selector *GovSelector `json:"selector,omitempty"`

	// 治理策略定义内容
	Spec           *interface{} `json:"spec,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListGovernancePolicyByPolicyIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGovernancePolicyByPolicyIdResponse struct{}"
	}

	return strings.Join([]string{"ListGovernancePolicyByPolicyIdResponse", string(data)}, " ")
}

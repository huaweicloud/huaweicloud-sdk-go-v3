package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGovernancePolicyResponseBody 治理策略详情
type ListGovernancePolicyResponseBody struct {
	MatchGroup *CreateBussinessScene `json:"matchGroup,omitempty"`

	// 治理策略定义
	Policies *[]ListGovernancePolicyResponsePolicies `json:"policies,omitempty"`
}

func (o ListGovernancePolicyResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGovernancePolicyResponseBody struct{}"
	}

	return strings.Join([]string{"ListGovernancePolicyResponseBody", string(data)}, " ")
}

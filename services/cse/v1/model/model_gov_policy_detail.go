package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GovPolicyDetail 治理策略详情。
type GovPolicyDetail struct {
	MatchGroup *CreateBussinessScene `json:"matchGroup,omitempty"`

	// 治理策略定义。
	Policies *[]GovPolicyDetailPolicies `json:"policies,omitempty"`
}

func (o GovPolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GovPolicyDetail struct{}"
	}

	return strings.Join([]string{"GovPolicyDetail", string(data)}, " ")
}

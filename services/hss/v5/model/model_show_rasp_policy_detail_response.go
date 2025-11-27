package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRaspPolicyDetailResponse Response Object
type ShowRaspPolicyDetailResponse struct {

	// 防护策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释** 操作系统类型 **取值范围** 包含以下两种: - Linux : linux系统 - Windows: windows系统
	OsType *string `json:"os_type,omitempty"`

	// list
	RuleList       *[]CheckFeatureRuleInfo `json:"rule_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowRaspPolicyDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRaspPolicyDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowRaspPolicyDetailResponse", string(data)}, " ")
}

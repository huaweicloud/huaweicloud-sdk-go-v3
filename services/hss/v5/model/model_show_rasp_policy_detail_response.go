package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRaspPolicyDetailResponse Response Object
type ShowRaspPolicyDetailResponse struct {

	// 防护策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 操作系统类型
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

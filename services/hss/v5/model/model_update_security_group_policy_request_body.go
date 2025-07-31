package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityGroupPolicyRequestBody struct {

	// 安全策略ID
	PolicyId string `json:"policy_id"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 安全组列表
	SecurityGroups []SecurityGroup `json:"security_groups"`
}

func (o UpdateSecurityGroupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupPolicyRequestBody", string(data)}, " ")
}

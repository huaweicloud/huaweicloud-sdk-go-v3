package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityGroupRuleRequest Request Object
type DeleteSecurityGroupRuleRequest struct {

	// 安全组规则ID。
	SecurityGroupRuleId string `json:"security_group_rule_id"`
}

func (o DeleteSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupRuleRequest", string(data)}, " ")
}

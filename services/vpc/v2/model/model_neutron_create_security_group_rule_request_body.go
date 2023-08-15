package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateSecurityGroupRuleRequestBody
type NeutronCreateSecurityGroupRuleRequestBody struct {
	SecurityGroupRule *NeutronCreateSecurityGroupRuleOption `json:"security_group_rule"`
}

func (o NeutronCreateSecurityGroupRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRuleRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRuleRequestBody", string(data)}, " ")
}

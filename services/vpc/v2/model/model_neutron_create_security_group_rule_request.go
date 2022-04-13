package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NeutronCreateSecurityGroupRuleRequest struct {
	Body *NeutronCreateSecurityGroupRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRuleRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComplianceRuleResponse Response Object
type DeleteComplianceRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteComplianceRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComplianceRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteComplianceRuleResponse", string(data)}, " ")
}

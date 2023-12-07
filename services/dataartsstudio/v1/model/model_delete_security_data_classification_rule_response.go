package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityDataClassificationRuleResponse Response Object
type DeleteSecurityDataClassificationRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityDataClassificationRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityDataClassificationRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityDataClassificationRuleResponse", string(data)}, " ")
}

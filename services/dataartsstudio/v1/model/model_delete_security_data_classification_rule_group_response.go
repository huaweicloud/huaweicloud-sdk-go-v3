package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityDataClassificationRuleGroupResponse Response Object
type DeleteSecurityDataClassificationRuleGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityDataClassificationRuleGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityDataClassificationRuleGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityDataClassificationRuleGroupResponse", string(data)}, " ")
}

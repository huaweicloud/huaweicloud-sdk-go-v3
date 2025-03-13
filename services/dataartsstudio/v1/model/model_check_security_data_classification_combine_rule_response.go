package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSecurityDataClassificationCombineRuleResponse Response Object
type CheckSecurityDataClassificationCombineRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckSecurityDataClassificationCombineRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSecurityDataClassificationCombineRuleResponse struct{}"
	}

	return strings.Join([]string{"CheckSecurityDataClassificationCombineRuleResponse", string(data)}, " ")
}

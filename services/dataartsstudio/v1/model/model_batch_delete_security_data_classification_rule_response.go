package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityDataClassificationRuleResponse Response Object
type BatchDeleteSecurityDataClassificationRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecurityDataClassificationRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityDataClassificationRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityDataClassificationRuleResponse", string(data)}, " ")
}

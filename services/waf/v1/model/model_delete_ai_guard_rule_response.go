package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiGuardRuleResponse Response Object
type DeleteAiGuardRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAiGuardRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiGuardRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAiGuardRuleResponse", string(data)}, " ")
}

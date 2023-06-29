package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopRuleResponse Response Object
type StopRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopRuleResponse struct{}"
	}

	return strings.Join([]string{"StopRuleResponse", string(data)}, " ")
}

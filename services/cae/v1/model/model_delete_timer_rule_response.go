package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTimerRuleResponse Response Object
type DeleteTimerRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTimerRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTimerRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteTimerRuleResponse", string(data)}, " ")
}

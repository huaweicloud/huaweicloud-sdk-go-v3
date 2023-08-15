package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTimerRuleResponse Response Object
type UpdateTimerRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTimerRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimerRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateTimerRuleResponse", string(data)}, " ")
}

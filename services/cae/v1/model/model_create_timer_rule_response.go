package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTimerRuleResponse Response Object
type CreateTimerRuleResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“TimerRule”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	Items          *[]TimerRuleDetails `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreateTimerRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTimerRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateTimerRuleResponse", string(data)}, " ")
}

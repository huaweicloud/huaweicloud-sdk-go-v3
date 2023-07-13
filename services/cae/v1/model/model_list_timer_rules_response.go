package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTimerRulesResponse Response Object
type ListTimerRulesResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“TimerRule”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	Items          *[]TimerRuleDetails `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListTimerRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTimerRulesResponse struct{}"
	}

	return strings.Join([]string{"ListTimerRulesResponse", string(data)}, " ")
}

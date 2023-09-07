package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTimerRulesResponse Response Object
type ListTimerRulesResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *TimeRuleKindObj `json:"kind,omitempty"`

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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTimerRuleResponse Response Object
type CreateTimerRuleResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *TimeRuleKindObj `json:"kind,omitempty"`

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

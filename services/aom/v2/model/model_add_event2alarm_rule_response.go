package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEvent2alarmRuleResponse Response Object
type AddEvent2alarmRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddEvent2alarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEvent2alarmRuleResponse struct{}"
	}

	return strings.Join([]string{"AddEvent2alarmRuleResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKeywordsAlarmRuleResponse Response Object
type CreateKeywordsAlarmRuleResponse struct {

	// 告警规则id
	KeywordsAlarmRuleId *string `json:"keywords_alarm_rule_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateKeywordsAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeywordsAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateKeywordsAlarmRuleResponse", string(data)}, " ")
}

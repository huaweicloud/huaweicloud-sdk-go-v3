package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteKeywordsAlarmRuleRequest struct {
	// 关键词告警规则id

	KeywordsAlarmRuleId string `json:"keywords_alarm_rule_id"`
}

func (o DeleteKeywordsAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeywordsAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeywordsAlarmRuleRequest", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListKeywordsAlarmRulesResponse struct {
	// 项目id

	KeywordsAlarmRules *[]KeywordsAlarmRuleRespList `json:"keywords_alarm_rules,omitempty"`
	HttpStatusCode     int                          `json:"-"`
}

func (o ListKeywordsAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeywordsAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListKeywordsAlarmRulesResponse", string(data)}, " ")
}

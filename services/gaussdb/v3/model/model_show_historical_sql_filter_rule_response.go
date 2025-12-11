package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoricalSqlFilterRuleResponse Response Object
type ShowHistoricalSqlFilterRuleResponse struct {

	// **参数解释**：  历史SQL限流规则。
	SqlFilterRules *[]HistoricalSqlFilterRule `json:"sql_filter_rules,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowHistoricalSqlFilterRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoricalSqlFilterRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoricalSqlFilterRuleResponse", string(data)}, " ")
}

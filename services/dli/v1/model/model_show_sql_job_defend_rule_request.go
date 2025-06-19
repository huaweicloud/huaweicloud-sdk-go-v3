package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobDefendRuleRequest Request Object
type ShowSqlJobDefendRuleRequest struct {

	// 拦截规则唯一标识。
	RuleId string `json:"rule_id"`
}

func (o ShowSqlJobDefendRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobDefendRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobDefendRuleRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlJobDefendRuleRequest Request Object
type DeleteSqlJobDefendRuleRequest struct {

	// 拦截规则唯一标识。
	RuleId string `json:"rule_id"`
}

func (o DeleteSqlJobDefendRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlJobDefendRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlJobDefendRuleRequest", string(data)}, " ")
}

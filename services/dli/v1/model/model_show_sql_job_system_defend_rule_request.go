package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlJobSystemDefendRuleRequest Request Object
type ShowSqlJobSystemDefendRuleRequest struct {

	// 系统规则唯一标识。
	RuleId string `json:"rule_id"`
}

func (o ShowSqlJobSystemDefendRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobSystemDefendRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlJobSystemDefendRuleRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlJobDefendRuleRequest Request Object
type UpdateSqlJobDefendRuleRequest struct {

	// 拦截规则唯一标识。
	RuleId string `json:"rule_id"`

	Body *SqlJobDefendRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateSqlJobDefendRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlJobDefendRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlJobDefendRuleRequest", string(data)}, " ")
}

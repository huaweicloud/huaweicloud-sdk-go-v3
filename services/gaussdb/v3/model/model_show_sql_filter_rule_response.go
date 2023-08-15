package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlFilterRuleResponse Response Object
type ShowSqlFilterRuleResponse struct {

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// SQL限流规则
	SqlFilterRules *[]SqlFilterRule `json:"sql_filter_rules,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowSqlFilterRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlFilterRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlFilterRuleResponse", string(data)}, " ")
}

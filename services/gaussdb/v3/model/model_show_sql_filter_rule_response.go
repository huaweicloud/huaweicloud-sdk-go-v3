package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSqlFilterRuleResponse struct {

	// 节点id
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// SQL限流规则
	SqlFilterRules *[]SqlFilterRule `json:"sql_filter_rules,omitempty" xml:"sql_filter_rules"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowSqlFilterRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlFilterRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlFilterRuleResponse", string(data)}, " ")
}

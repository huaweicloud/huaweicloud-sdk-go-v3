package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSqlFilterRule SQL限流规则。
type NodeSqlFilterRule struct {

	// Sql限流类型。  取值范围： - SELECT - UPDATE - DELETE
	SqlType string `json:"sql_type"`

	// SQL限流具体规则。
	Patterns []NodeSqlFilterRulePattern `json:"patterns"`
}

func (o NodeSqlFilterRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSqlFilterRule struct{}"
	}

	return strings.Join([]string{"NodeSqlFilterRule", string(data)}, " ")
}

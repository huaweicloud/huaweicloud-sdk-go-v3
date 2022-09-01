package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SQL限流规则。
type SqlFilterRule struct {

	// Sql限流类型。  取值范围： - SELECT - UPDATE - DELETE
	SqlType string `json:"sql_type" xml:"sql_type"`

	// SQL限流具体规则。
	Patterns []SqlFilterRulePattern `json:"patterns" xml:"patterns"`
}

func (o SqlFilterRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlFilterRule struct{}"
	}

	return strings.Join([]string{"SqlFilterRule", string(data)}, " ")
}

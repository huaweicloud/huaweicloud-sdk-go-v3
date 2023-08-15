package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodeSqlFilterRule SQL限流规则。
type DeleteNodeSqlFilterRule struct {

	// Sql限流类型。  取值范围： - SELECT - UPDATE - DELETE
	SqlType string `json:"sql_type"`

	// SQL限流具体规则。
	Patterns []string `json:"patterns"`
}

func (o DeleteNodeSqlFilterRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeSqlFilterRule struct{}"
	}

	return strings.Join([]string{"DeleteNodeSqlFilterRule", string(data)}, " ")
}

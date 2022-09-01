package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// - 开启/关闭sql限流参数体。
type OperateSqlFilterRuleReq struct {
	SqlFilterRules []NodeSqlFilterRuleInfo `json:"sql_filter_rules" xml:"sql_filter_rules"`
}

func (o OperateSqlFilterRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateSqlFilterRuleReq struct{}"
	}

	return strings.Join([]string{"OperateSqlFilterRuleReq", string(data)}, " ")
}

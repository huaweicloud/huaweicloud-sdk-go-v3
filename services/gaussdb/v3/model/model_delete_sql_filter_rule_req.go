package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlFilterRuleReq - 开启/关闭sql限流参数体。
type DeleteSqlFilterRuleReq struct {
	SqlFilterRules []DeleteNodeSqlFilterRuleInfo `json:"sql_filter_rules"`
}

func (o DeleteSqlFilterRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlFilterRuleReq struct{}"
	}

	return strings.Join([]string{"DeleteSqlFilterRuleReq", string(data)}, " ")
}

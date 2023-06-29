package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSqlFilterRuleInfo 节点级别的SQL限流规则。
type NodeSqlFilterRuleInfo struct {

	// 节点ID
	NodeId string `json:"node_id"`

	// SQL限流规则。集合的sql_type值不能重复。
	Rules []NodeSqlFilterRule `json:"rules"`
}

func (o NodeSqlFilterRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSqlFilterRuleInfo struct{}"
	}

	return strings.Join([]string{"NodeSqlFilterRuleInfo", string(data)}, " ")
}

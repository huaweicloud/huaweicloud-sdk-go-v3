package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodeSqlFilterRuleInfo 节点级别的SQL限流规则。
type DeleteNodeSqlFilterRuleInfo struct {

	// 节点ID
	NodeId string `json:"node_id"`

	// SQL限流规则。
	Rules []DeleteNodeSqlFilterRule `json:"rules"`
}

func (o DeleteNodeSqlFilterRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeSqlFilterRuleInfo struct{}"
	}

	return strings.Join([]string{"DeleteNodeSqlFilterRuleInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点级别的SQL限流规则。
type DeleteNodeSqlFilterRuleInfo struct {

	// 节点id
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LimitTaskRuleInfoOption struct {

	// 节点id。
	NodeId string `json:"node_id"`

	// 规则id。
	RuleId string `json:"rule_id"`

	// 状态，当前支持：CREATING，UPDATEING，DELETING，WAIT_EXCUTE，EXCUTING，TIME_OVER，DELETED，CREATE_FAILED，UPDATE_FAILED，DELETE_FAILED，EXCEPTION。
	Status string `json:"status"`
}

func (o LimitTaskRuleInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitTaskRuleInfoOption struct{}"
	}

	return strings.Join([]string{"LimitTaskRuleInfoOption", string(data)}, " ")
}

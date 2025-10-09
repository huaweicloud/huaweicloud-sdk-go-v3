package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateHtapQueryQueueRuleReq **参数解释**：  设置查询当前查询队列阈值请求体。  **约束限制**：  不涉及。
type OperateHtapQueryQueueRuleReq struct {
	QueryQueueRule *HtapQueryQueueRule `json:"query_queue_rule"`
}

func (o OperateHtapQueryQueueRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateHtapQueryQueueRuleReq struct{}"
	}

	return strings.Join([]string{"OperateHtapQueryQueueRuleReq", string(data)}, " ")
}

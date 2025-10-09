package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateHtapQueryQueueControlReq **参数解释**：  开启/关闭查询队列参数体。  **约束限制**：  不涉及。
type OperateHtapQueryQueueControlReq struct {

	// **参数解释**：  查询队列的开关状态。  **约束限制**：  不涉及。  **取值范围**：  - true：表示开启。 - false：表示关闭。  **默认值**：  false。
	EnableQueryQueueSelect string `json:"enable_query_queue_select"`
}

func (o OperateHtapQueryQueueControlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateHtapQueryQueueControlReq struct{}"
	}

	return strings.Join([]string{"OperateHtapQueryQueueControlReq", string(data)}, " ")
}
